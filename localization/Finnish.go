package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Finnish] = map[string]string{
		"#admin": "järjestelmänvalvoja",
		"#admin.allow-role": "salli-roolin",
		"#look": "katso",
		"#move": "liikkua",
		"#review": "arvostelu",
		"#admin.allow-role.value": "arvo",
		"#.subject": "kohde",
		"#look.info": "tiedot",
		"#look.review-list": "arvosteluluettelo",
		"#admin.allow-role.Description": "Näytä tulokset rooleina. Oletus: False",
		"#look.Description": "nähdä jotain",
		"#move.Description": "Siirrä arvostelu tälle kanavalle",
		"#review.Description": "Käyttäjän arvostelu",
		"#admin.allow-role.value.Description": "aseta arvo",
		"#.subject.Description": "Valitse kohde",
		"#look.info.Description": "Katso käyttäjätiedot",
		"#look.review-list.Description": "Katso luettelo käyttäjien saamista arvosteluista",
		"#allow-role.NeedPermissions": "Botilla ei ole käyttöoikeuksia - Hallitse rooleja",
		"#allow-role.InProgress": "Työn alla",
		"#allow-role.proc.Title": "Muokkaa asetuksia",
		"#allow-role.proc.Description": "Salli rooli",
		"#allow-role.proc.InProgress": "Jatketaan",
		"#allow-role.proc.Done": "saattaa loppuun",
		"#allow-role.Keep": "Ei asetuksia muutettu",
		"#look.info.IsNone": "Ei arvosteluja",
		"#look.review-list.IsNone": "Ei arvosteluja",
		"#look.review-list.menu.Title": "Arvostelut kohteelle %s",
		"#look.review-list.menu.Page": "%d sivu",
		"#move.IsNone": "Kohteelle ei ole kirjoitettu arvosteluja",
		"#move.Move": "muuta tänne",
		"#review.SelfReview": "Et voi arvostella itseäsi",
		"#review.modal.Title": "Arvostelut",
		"#review.lable.Score": "pisteet",
		"#review.lable.Title": "otsikko",
		"#review.lable.Content": "yksityiskohta",
		"$review.IsEdited": "Tätä arvostelua on muokattu",
		"$review.NoAuthor": "Poistettuja arvosteluja ei voi palauttaa, koska kirjoittajaa ei ole täällä",
		"$review.DM": "Uusi arvostelu on kirjoitettu",
    }
}
