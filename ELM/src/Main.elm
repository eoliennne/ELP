module Main exposing (..)


import Browser
import Html exposing (..)
import Html.Events exposing (onClick, onInput)
import Html.Attributes exposing (placeholder,value)
import Http
import Json.Decode exposing (Decoder, list, map2, string, int, field,at)
import Random exposing (Generator,generate,constant,int)
import List.Extra exposing (getAt)

--FONCTIONS

verifSol : String -> Model -> Bool
verifSol str model = 
    if str==model.current_word.word then True else False

getWords : Cmd Msg
getWords =
    Http.get
        { url = "http://localhost:5018/thousand_words_things_explainer.txt"
        , expect = Http.expectString GotList
        }

getRandomWord : List String -> Random.Generator (Maybe String)
getRandomWord listeMots =
    let
        indexMax = List.length listeMots
        generateurIndex = Random.int 0 (indexMax - 1)
    in
    if indexMax == 0 then
        Random.constant Nothing
    else
        Random.map (\index -> getAt index listeMots) generateurIndex

convertToWord : Maybe String -> Word
convertToWord mot = 
    case mot of
        Just str -> {word = str, meanings = [{wordtype = "type", definition = ["def2"]}]}
        Nothing -> {word = "word", meanings = [{wordtype = "type", definition = ["def2"]}]}

-- DECODERS

firstDecoder = at ["0"](categoryDecoder)

categoryDecoder : Decoder Word
categoryDecoder =
    map2 Word
        (field "word" string)
        (field "meanings" (list meaningsDecoder))

meaningsDecoder : Decoder Meanings
meaningsDecoder =
    map2 Meanings
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

definitionDecoder : Decoder String
definitionDecoder = 
    field "definition" string
    
viewMeanings : Meanings -> Html Msg
viewMeanings meanings =
  div []
    [ div [] [ h3 [] [text  meanings.wordtype] ]
    , ul [] (List.map (\definition -> li [] [ text definition ]) meanings.definition)
    ]
    
errorToString : Http.Error -> String
errorToString error =
    case error of
        Http.BadUrl url ->
            "The URL " ++ url ++ " was invalid"
        Http.Timeout ->
            "Unable to reach the server, try again"
        Http.NetworkError ->
            "Unable to reach the server, check your network connection"
        Http.BadStatus 500 ->
            "The server had a problem, try again later"
        Http.BadStatus 400 ->
            "Verify your information and try again"
        Http.BadStatus x ->
            "Unknown error with status " ++ (String.fromInt x)
        Http.BadBody errorMessage ->
            errorMessage


--TYPES
type alias Meanings = {wordtype:String, definition:(List String)}
type alias Word = {word:String,meanings:(List Meanings)}


--MAIN
main : Program () Model Msg
main =
  Browser.element { init = init, update = update, view = view, subscriptions = subscriptions }



--MODEL

type alias Model = {current_word : Word, solution : String, statut : Status, load : LoadWord}
type Status = Right | Wrong | NoSol
type LoadWord
  = Failure String
  | Loading
  | Success Word

firstModel : Model
firstModel = Model {word = "word", meanings = [{wordtype = "type", definition = ["def1"]}]} "" NoSol Loading


--INIT
init : () -> (Model, Cmd Msg)
init _ = (firstModel 
    , Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ firstModel.current_word.word
      , expect = Http.expectJson GotData firstDecoder
      })



--UPDATE
type Msg = GetSol 
        | AskWord 
        | CheckAns String 
        | GotData (Result Http.Error Word)
        | GotList (Result Http.Error String)
        | GotNewWord (Maybe String)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model  =
  case msg of
    CheckAns newContent-> ({model | solution = newContent}, Cmd.none)
    GetSol ->  if (verifSol model.solution model) then ({model | statut = Right},Cmd.none)
                    else ({model | statut = Wrong}, Cmd.none)
    GotNewWord newWord -> ({model | current_word = convertToWord newWord, statut = NoSol}, Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ (convertToWord newWord).word
      , expect = Http.expectJson GotData firstDecoder
      })
    GotData result -> case result of
          Ok data ->
            ({model | load = Success data}, Cmd.none)

          Err probleme ->
            ({model | load = Failure (errorToString probleme)}, Cmd.none)

    AskWord -> (model,getWords)
    GotList (Ok data) ->
            let
                words = String.split " " data 
                generator = getRandomWord words
            in
            ( model, Random.generate GotNewWord generator )
    GotList (Err _) ->
        ( model, Cmd.none )


--VIEW

view : Model -> Html Msg
view model =
         div []
        [ h1 [] [text "Guess It !"]
        , viewDef model
        , viewSol model
        ,viewField model
        ]

viewDef : Model -> Html Msg
viewDef model =
  case model.load of
    Failure message->
      text message

    Loading ->
      text "Loading definition..."

    Success chosenword ->
      div []
        [ --div [] [ text ("Word: " ++ chosenword.word) ]
         h2 [] [ text "Meanings" ]
        , ul [] (List.map viewMeanings chosenword.meanings)
        ]

viewField : Model -> Html Msg
viewField model = div []
        [
        div [] [input [placeholder "Enter a word", value model.solution, onInput CheckAns] []]
        , div [] [button [onClick GetSol] [text "Solution"]]
        , div [] [button [onClick AskWord] [text "Refresh"]]
         ]

viewSol : Model -> Html Msg
viewSol model = 
    case model.statut of
        Right -> div[]
            [div[] [text "That is correct !"]]
        Wrong -> div[]
            [div[] [text "Incorrect. The answer is : "]
            , text model.current_word.word]
        NoSol -> div [] [text ""]



--SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none