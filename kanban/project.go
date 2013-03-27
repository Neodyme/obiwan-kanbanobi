package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Project struct {
	Id        uint32
	Name      string
	admins_id []uint32
	Read      []uint32
	Content   string
}

func MsgProjectCreate(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		0,
		*msg.Projects.Name,
		msg.Projects.AdminsId,
		msg.Projects.Read,
		"",
	}
	var answer *message.Msg
	if err := proj.Add(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(31), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectCreate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectUpdate(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		0,
		*msg.Projects.Name,
		msg.Projects.AdminsId,
		msg.Projects.Read,
		"",
	}
	var answer *message.Msg
	if err := proj.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(32), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectDelete(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Id: *msg.Projects.Id,
	}
	var answer *message.Msg
	if err := proj.Del(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(34), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectGet(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Id: *msg.Projects.Id,
	}
	var answer *message.Msg
	if err := proj.Get(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(35), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Projects: &message.Msg_Projects{

				Id:   proto.Uint32(proj.Id),
				Name: proto.String(proj.Name),
				//				AdminsId:   proto.Uint32(proj.admins_id),
				//		:    "",
			},
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectGetBoard(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		*msg.Projects.Id,
		*msg.Projects.Name,
		msg.Projects.AdminsId,
		msg.Projects.Read,
		"",
	}
        var answer *message.Msg
	
//        if board, err := user.GetProjectByUserId(dbPool); err != nil{
                answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
                Error: &message.Msg_Error{
				ErrorId: proto.Uint32(15),
			},
		}
//	} else {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		Users: &message.Msg_Cards{
				Id:          proto.Uint32(proj.Id),
	Name:        &proj.Name,
				Admin:       &proj.Admin,
				UserProject: ConvertTabOfCardToMessage(board),
			},
		}
//	}
	sendKanbanMsg(conn, answer)
}

// modifier pour faire des column
func ConvertTabOfColumnToMessage(p []Column) []*message.Msg_Column {
        var ret []*message.Msg_Column
	
        for n := 0; n < len(p); n++ {
                ret = append(ret, &message.Msg_Cards{
			Id:         proto.Uint32(p[n].Id),
			ProjectId:  proto.Uint32(p[n].Project_id),
			ColumnId:   proto.Uint32(p[n].Column_id),
			Name:       proto.String(p[n].Name),
			Desc:       proto.String(p[n].Content),
			Tags:       p[n].Tags,
			UserId:     proto.Uint32(p[n].User_id),
			ScriptsIds: p[n].Scripts_id,
			Write:      p[n].Write,
                })
        }
        return ret
}

// Cette fonction a une gestion synchrone des messages (traitement les uns apres les autres, pas de traitements paralleles)
// Il faut faire une pool de worker, un dispacher et lancer l'operation a effectuer dans le dispatch.
func MsgProject(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgProjectCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgProjectUpdate(conn, msg)
	case message.CMD_DELETE:
		MsgProjectDelete(conn, msg)
	case message.CMD_GET:
		MsgProjectGet(conn, msg)
	case message.CMD_MOVE:
		MsgProjectUpdate(conn, msg)
	case message.CMD_GETBOARD:
		MsgProjectGetBoard(conn, msg)
        default:
                UnknowCommand(conn, msg)
	}
}
