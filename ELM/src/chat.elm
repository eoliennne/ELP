module Main exposing (..)

import Browser exposing (Browser, Html, Url, div, button, text)
import Html exposing (Html)
import Http exposing (Request, Response)
import Json.Decode exposing (Decoder, string, succeed)

type Msg
    = FetchRandomWord
    | GotRandomWord (Result Http.Error String)
    | FetchInfo
    | GotInfo (Result Http.Error String)

type Model =
    { randomWord : Maybe String
    , info : Maybe String
    }

init : Model
init =
    { randomWord = Nothing
    , info = Nothing
    }

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        FetchRandomWord ->
            (model, Http.get { url = "https://api.example.com/random-word", expect = Http.expectJson GotRandomWord wordDecoder })

        GotRandomWord (Ok word) ->
            ( { model | randomWord = Just word }, Cmd.none )

        GotRandomWord (Err _) ->
            (model, Cmd.none)

        FetchInfo ->
            case model.randomWord of
                Just word ->
                    (model, Http.get { url = "https://api.example.com/info/" ++ word, expect = Http.expectString GotInfo })

                Nothing ->
                    (model, Cmd.none)

        GotInfo (Ok info) ->
            ( { model | info = Just info }, Cmd.none )

        GotInfo (Err _) ->
            (model, Cmd.none)

view : Model -> Html Msg
view model =
    div []
        [ button [ onClick FetchRandomWord ] [ text "Fetch Random Word" ]
        , button [ onClick FetchInfo ] [ text "Fetch Info" ]
        , div []
            [ case model.randomWord of
                Just word ->
                    text ("Random Word: " ++ word)

                Nothing ->
                    text "No random word yet."
            ]
        , div []
            [ case model.info of
                Just info ->
                    text ("Info: " ++ info)

                Nothing ->
                    text "No info yet."
            ]
        ]

main : Program () Model Msg
main =
    Browser.sandbox { init = init, update = update, view = view, subscriptions = \_ -> Sub.none }

wordDecoder : Decoder String
wordDecoder =
    string
