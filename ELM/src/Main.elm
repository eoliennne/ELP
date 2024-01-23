module Main exposing (..)

--Test structure

import Browser
import Html exposing (..)
import Html.Events exposing (onClick, onInput)
import Html.Attributes exposing (placeholder,value,style)
import Http
import Json.Decode exposing (Decoder, list, map2, string, int, field,at)

--FONCTIONS

verifSol : String -> Model -> Bool
verifSol str model = 
    if str==model.current_word.word then True else False

    --a changer : doit retourner un nouveau mot
new_word : Word
new_word = {word = "Help", meanings = [{wordtype = "type", definition = ["def2"]}]}



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
type Msg = GetSol | GetNewWord | CheckAns String | GotData (Result Http.Error Word)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model  =
  case msg of
    CheckAns newContent-> ({model | solution = newContent}, Cmd.none)
    GetSol ->  if (verifSol model.solution model) then ({model | statut = Right},Cmd.none)
                    else ({model | statut = Wrong}, Cmd.none)
    GetNewWord  -> ({model | current_word = new_word, statut = NoSol}, Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ model.current_word.word
      , expect = Http.expectJson GotData firstDecoder
      })
    GotData result -> case result of
          Ok data ->
            ({model | load = Success data}, Cmd.none)

          Err probleme ->
            ({model | load = Failure (errorToString probleme)}, Cmd.none)


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
        , div [] [button [onClick GetNewWord] [text "Refresh"]]
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
subscriptions model =
    Sub.none