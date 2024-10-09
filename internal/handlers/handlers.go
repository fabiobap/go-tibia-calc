package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fabiobap/go-tibia-calc/helpers"
	"github.com/fabiobap/go-tibia-calc/internal/config"
	"github.com/fabiobap/go-tibia-calc/internal/forms"
	"github.com/fabiobap/go-tibia-calc/internal/models"
	"github.com/fabiobap/go-tibia-calc/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

type InfoLevelResponse struct {
	Hitpoints        int    `json:"hitpoints"`
	Manapoints       int    `json:"manapoints"`
	Cap              int    `json:"cap"`
	Experience       int    `json:"experience"`
	OneRegularBless  string `json:"one_reg_bless"`
	FiveRegularBless string `json:"five_reg_bless"`
	TwistBless       string `json:"twist_bless"`
	SevenBless       string `json:"seven_bless"`
	FullBless        string `json:"full_bless"`
}
type MidnightShardsResponse struct {
	Experience int `json:"experience"`
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) InfoLevel(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["vocation"] = "none"

	intMap := make(map[string]int)
	intMap["level"] = 1

	render.Template(w, r, "info-lvl.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		IntMap:    intMap,
		StringMap: stringMap,
	})
}

func (m *Repository) PostInfoLevel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("level", "vocation")
	form.Minlength("level", 1)

	level, _ := strconv.Atoi(r.Form.Get("level"))
	vocation := r.Form.Get("vocation")

	character := models.Character{
		Vocation: vocation,
		Level:    level,
	}

	character.Load()

	if !form.Valid() {
		render.Template(w, r, "info-lvl.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	resp := InfoLevelResponse{
		Hitpoints:        character.Hitpoints,
		Manapoints:       character.Manapoints,
		Cap:              character.Cap,
		Experience:       character.Experience,
		OneRegularBless:  render.Gold(character.BlessingRegularOne),
		TwistBless:       render.Gold(character.BlessingTwist),
		FiveRegularBless: render.Gold(character.BlessingRegularFive),
		SevenBless:       render.Gold(character.BlessingSeven),
		FullBless:        render.Gold(character.BlessingFull),
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")

	w.Write(out)
}

func (m *Repository) MidnightShards(w http.ResponseWriter, r *http.Request) {
	intMap := make(map[string]int)
	intMap["level"] = 1
	intMap["quantity"] = 1

	render.Template(w, r, "midnight-shards.page.tmpl", &models.TemplateData{
		Form:   forms.New(nil),
		IntMap: intMap,
	})
}

func (m *Repository) PostMidnightShards(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("level", "qty")
	form.Minlength("level", 1)
	form.Minlength("qty", 1)

	level, _ := strconv.Atoi(r.Form.Get("level"))
	qty, _ := strconv.Atoi(r.Form.Get("qty"))

	ms := models.MidnightShard{
		Quantity: qty,
		Level:    level,
	}

	ms.Load()

	if !form.Valid() {
		render.Template(w, r, "midnight-shards.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	resp := MidnightShardsResponse{
		Experience: ms.Experience,
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")

	w.Write(out)
}

func (m *Repository) StoneOfInsight(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "soi.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostStoneOfInsight(w http.ResponseWriter, r *http.Request) {
}
