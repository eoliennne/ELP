import Json.Decode as D
import Http
import Html exposing (Html, button, div, text,pre)
import Browser

--TYPES
-- type stockant la nature et les définitions correspondantes d'un mot
type alias Definition = {wordtype:String, meaning:(List String)}
-- Registre stockant un mot et toutes ses définitions         
type alias Word = {word:String,definition:Definition}



--MAIN

first = "potato"

main =
  Browser.element
    { init = init first
    , update = update
    , subscriptions = subscriptions
    , view = view
    }

-- MODEL

type Model
  = Failure
  | Loading
  | Success String


init : String -> () -> (Model, Cmd Msg)
init firstword _ =
  ( Loading
  , Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ firstword
      , expect = Http.expectString GotData
      }
  )

-- UPDATE

type Msg
  = GotData (Result Http.Error String)

update msg model =
  case msg of
    GotData result ->
      case result of
          Ok data ->
            (Success data, Cmd.none)

          Err _ ->
            (Failure, Cmd.none)

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none

-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure ->
      text "I was unable to load the JSON."

    Loading ->
      text "Loading..."

    Success fullText ->
      pre [] [ text fullText ]

wordDecoder =
  map2 Word
    (field "word" string)
    (field "source" string)
    (field "author" string)
    (field "year" int)
