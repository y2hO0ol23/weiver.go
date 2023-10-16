package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Polish] = map[string]string{
		"#admin": "admin",
		"#admin.allow-role": "pozwolić-na-rolę",
		"#look": "patrzeć",
		"#move": "przenosić",
		"#review": "recenzja",
		"#admin.allow-role.value": "wartość",
		"#.subject": "cel",
		"#look.info": "informacja",
		"#look.review-list": "lista-recenzji",
		"#admin.allow-role.Description": "Wyświetlaj wyniki jako role. Wartość domyślna: Fałsz",
		"#look.Description": "zobaczyć coś",
		"#move.Description": "Przenieś recenzję na ten kanał",
		"#review.Description": "Opinia użytkownika",
		"#admin.allow-role.value.Description": "ustalić wartość",
		"#.subject.Description": "Wybierz cel",
		"#look.info.Description": "Wyświetl informacje o użytkowniku",
		"#look.review-list.Description": "Zobacz listę recenzji otrzymanych przez użytkowników",
		"#allow-role.NeedPermissions": "Bot nie ma uprawnień - Zarządzaj rolami",
		"#allow-role.InProgress": "Praca w toku",
		"#allow-role.proc.Title": "Zmień opcje",
		"#allow-role.proc.Description": "Zezwalaj na rolę",
		"#allow-role.proc.InProgress": "Postępowanie",
		"#allow-role.proc.Done": "kompletny",
		"#allow-role.Keep": "Żadne ustawienia nie uległy zmianie",
		"#look.info.IsNone": "Brak recenzji",
		"#look.review-list.IsNone": "Brak recenzji",
		"#look.review-list.menu.Title": "Recenzje %s",
		"#look.review-list.menu.Page": "%d strona",
		"#move.IsNone": "Dla celu nie napisano żadnych recenzji",
		"#move.Move": "przesuń się tutaj",
		"#review.SelfReview": "Nie możesz sam dokonać przeglądu",
		"#review.modal.Title": "Opinie",
		"#review.lable.Score": "wynik",
		"#review.lable.Title": "tytuł",
		"#review.lable.Content": "Szczegół",
		"$review.IsEdited": "Ta recenzja została edytowana",
		"$review.NoAuthor": "Usuniętych recenzji nie można przywrócić, ponieważ autor tutaj nie istnieje",
		"$review.DM": "Napisano nową recenzję",
    }
}
