module Main exposing (..)

--Test structure

import Browser
import Html exposing (..)
import Html.Events exposing (onClick)
import Html.Attributes exposing (placeholder,value)
import Html.Events exposing (onInput)

--FONCTIONS

verifSol : String -> Model -> Bool
verifSol str model = 
    if str==model.current_word.word then True else False

    --a changer : doit retourner un nouveau mot
gotWord : Word
gotWord = {word = "Help", meanings = [{wordtype = "type", definition = ["def2"]}]}


--TYPES
type alias Meanings = {wordtype:String, definition:(List String)}
type alias Word = {word:String,meanings:(List Meanings)}

--MAIN
main : Program () Model Msg
main =
  Browser.element { init = init, update = update, view = view, subscriptions = subscriptions }




--MODEL

type alias Model = {current_word : Word, solution : String, statut : Status}
type Status = Right | Wrong | NoSol

init : () -> (Model, Cmd Msg)
init _ = ({current_word = {word = "Premier", meanings = [{wordtype = "type", definition = ["def1"]}]}, solution = "", statut = NoSol}, Cmd.none)

type Msg = GetSol | GetNewWord | Change String


--UPDATE


update : Msg -> Model -> (Model, Cmd Msg)
update msg model  =
  case msg of
    Change newContent-> ({model | solution = newContent}, Cmd.none)
    GetSol ->  if (verifSol model.solution model) then ({model | statut = Right},Cmd.none)
                    else ({model | statut = Wrong}, Cmd.none)
    GetNewWord  -> ({model | current_word = gotWord, statut = NoSol}, Cmd.none)



--VIEW

view : Model -> Html Msg
view model =
         div []
        [ h1 [] [text "Guess It !"]
        , h2 [] [text "Meaning"]
        , div [] [text model.current_word.word] --affiche le mot pour l'instant
        , div [] [input [placeholder "Enter a word", value model.solution, onInput Change] []]
        , div [] [button [onClick GetSol] [text "Solution"]]
        , div [] [button [onClick GetNewWord] [text "Refresh"]]
        , viewSol model
        ]


viewSol : Model -> Html Msg
viewSol model = 
    case model.statut of
        Right -> div[]
            [div[] [text "That is Correct !"]
            ]
        Wrong -> div[]
            [div[] [text "Incorrect. The answer is : "]
            , text model.current_word.word
            ]
        NoSol -> div [] [text ""]



--SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none