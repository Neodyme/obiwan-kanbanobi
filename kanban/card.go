package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Card struct {
	Id         uint32
	Name       string
	Content    string
	Column_id  uint32
	Project_id uint32
	Tags       []string
	User_id    uint32
	Scripts_id []uint32
	Write      []uint32
}

// Il y a moyen de factoriser beaucoup le code des fonctions ici.
// Il faut juste penser a faire une gestion d'erreur un peu avant la fonction
// commune pour renvoyer le bon code d'erreur (verifier que l'ID d'une carte
// existe avant de faire un delete par exemple)

// msg.Cards.UserId est utilise par defaut pour le moment. Mais c'est un champ optionnel.
// Il faudrait faire un test pour savoir si c'est le author_id ou lui qui est utilise.
func MsgCardCreate(conn net.Conn, msg *message.Msg) {
	card := &Card{
		0,
		*msg.Cards.Name,
		*msg.Cards.Desc,
		*msg.Cards.ColumnId,
		*msg.Cards.ProjectId,
		msg.Cards.Tags,
		*msg.Cards.UserId,
		msg.Cards.ScriptsIds,
		msg.Cards.Write,
	}
	var answer *message.Msg
	proj := &Project{
		Id: *msg.Cards.ProjectId,
	}
	if adm, err := proj.IsAdmin(dbPool, *msg.AuthorId); adm == false || err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := card.Add(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(5), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		card.GetLastCardWithName(dbPool)
		println(card.Id)
		*msg.Cards.Id = card.Id
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Cards: &message.Msg_Cards{
				Id:         proto.Uint32(card.Id),
				ProjectId:  proto.Uint32(card.Project_id),
				ColumnId:   proto.Uint32(card.Column_id),
				Name:       proto.String(card.Name),
				Desc:       proto.String(card.Content),
				Tags:       card.Tags,
				UserId:     proto.Uint32(card.User_id),
				ScriptsIds: card.Scripts_id,
				Write:      card.Write,
			},
		}
		notifyUsers(msg)
	}
	sendKanbanMsg(conn, answer)
}

func MsgCardUpdate(conn net.Conn, msg *message.Msg) {
	card := &Card{
		*msg.Cards.Id,
		*msg.Cards.Name,
		*msg.Cards.Desc,
		*msg.Cards.ColumnId,
		*msg.Cards.ProjectId,
		msg.Cards.Tags,
		*msg.Cards.UserId,
		msg.Cards.ScriptsIds,
		msg.Cards.Write,
	}
	var answer *message.Msg
	proj := &Project{
		Id: *msg.Cards.ProjectId,
	}
	if adm, err := proj.IsAdmin(dbPool, *msg.AuthorId); adm == false || err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := card.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici

		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(6), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Cards: &message.Msg_Cards{
				Id:         proto.Uint32(card.Id),
				ProjectId:  proto.Uint32(card.Project_id),
				ColumnId:   proto.Uint32(card.Column_id),
				Name:       proto.String(card.Name),
				Desc:       proto.String(card.Content),
				Tags:       card.Tags,
				UserId:     proto.Uint32(card.User_id),
				ScriptsIds: card.Scripts_id,
				Write:      card.Write,
			},
		}
		notifyUsers(msg)
	}
	sendKanbanMsg(conn, answer)
}

func MsgCardDelete(conn net.Conn, msg *message.Msg) {
	card := &Card{
		Id: *msg.Cards.Id,
	}
	var answer *message.Msg
	proj := &Project{
		Id: *msg.Cards.ProjectId,
	}
	if adm, err := proj.IsAdmin(dbPool, *msg.AuthorId); adm == false || err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := card.Del(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(7), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
		notifyUsers(msg)
	}
	sendKanbanMsg(conn, answer)
}

func MsgCardGet(conn net.Conn, msg *message.Msg) {
	card := &Card{
		Id: *msg.Cards.Id,
	}
	var answer *message.Msg
	proj := &Project{
		Id: *msg.Cards.ProjectId,
	}
	if adm, err := proj.IsAdmin(dbPool, *msg.AuthorId); adm == false || err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := card.Get(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(8), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Cards: &message.Msg_Cards{
				Id:         proto.Uint32(card.Id),
				ProjectId:  proto.Uint32(card.Project_id),
				ColumnId:   proto.Uint32(card.Column_id),
				Name:       proto.String(card.Name),
				Desc:       proto.String(card.Content),
				Tags:       card.Tags,
				UserId:     proto.Uint32(card.User_id),
				ScriptsIds: card.Scripts_id,
				Write:      card.Write,
			},
		}
	}
	sendKanbanMsg(conn, answer)
}

// Cette fonction a une gestion synchrone des messages (traitement les uns apres les autres, pas de traitements paralleles)
// Il faut faire une pool de worker, un dispacher et lancer l'operation a effectuer dans le dispatch.
func MsgCard(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgCardCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgCardUpdate(conn, msg)
	case message.CMD_DELETE:
		MsgCardDelete(conn, msg)
	case message.CMD_GET:
		MsgCardGet(conn, msg)
	case message.CMD_MOVE:
		MsgCardUpdate(conn, msg)
	default:
		UnknowCommand(conn, msg)
	}
}
