module Load_random exposing (..)

import Json.Decode exposing (Decoder, list, map2, string, int, field,at)
import Http
import Html exposing (Html, button, div, text,pre,ul,li)
import Browser
import Random.List exposing (choose)

--MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }

-- MODEL

type Model
  = Failure String
  | Loading
  | Success (List String)


init : () -> (Model, Cmd Msg)
init  _ =
  ( Loading
  , Http.get
      { url = "../static/thousand_words.json" 
      , expect = Http.expectJson GotData firstDecoder
      }
  )

-- UPDATE

type Msg
  = GotData (Result Http.Error (List String))

update msg model =
  case msg of
    GotData result ->
      case result of
          Ok data ->
            (Success data, Cmd.none)

          Err probleme ->
            (Failure (errorToString probleme), Cmd.none)

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

    Success liste ->   
      pre [] [text (let (mot,nul) = (choose liste) in { model | hasard = mot })]


--FUNCTIONS
firstDecoder : Decoder (List String)
firstDecoder = at ["words"] (list string)

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
