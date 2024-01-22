import Json.Decode exposing (Decoder, list, map2, string, int, field,at)
import Http
import Html exposing (Html, button, div, text,pre,ul,li)
import Browser

--TYPES

type alias Meanings = {wordtype:String, definition:(List String)}
       
type alias Word = {word:String,meanings:(List Meanings)}


--MAIN

first = "word"

main =
  Browser.element
    { init = init first
    , update = update
    , subscriptions = subscriptions
    , view = view
    }

-- MODEL

type Model
  = Failure String
  | Loading
  | Success Word


init : String -> () -> (Model, Cmd Msg)
init firstword _ =
  ( Loading
  , Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ firstword
      , expect = Http.expectJson GotData firstDecoder
      }
  )

-- UPDATE

type Msg
  = GotData (Result Http.Error Word)

update msg model =
  case msg of
    GotData result ->
      case result of
          Ok data ->
            (Success data, Cmd.none)

          Err truc ->
            (Failure (errorToString truc), Cmd.none)

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none

-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure message->
      text message

    Loading ->
      text "Loading..."

    Success chosenword ->
      div []
        [ pre [] [ text "Word loaded!" ]
        , div [] [ text ("Word: " ++ chosenword.word) ]
        , div [] [ text "Meanings:" ]
        , ul [] (List.map viewMeanings chosenword.meanings)
        ]


--FUNCTIONS

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
    [ div [] [ text ("Part of Speech: " ++ meanings.wordtype) ]
    , div [] [ text "Definitions:" ]
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
