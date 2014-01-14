// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

//	MeetmeList
//		Lists all users in a particular MeetMe conference. 
//		MeetmeList will follow as separate events, followed by a final event called MeetmeListComplete.	
//
func MeetmeList(socket *Socket, actionID, conference string) ([]map[string]string, error) {
	if len(actionID) == 0 || len(conference) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: MeetmeList",
		"\r\nActionID: ",
		actionID,
		"\r\nConference: ",
		conference,
		"\r\n\r\n", // end of command
	}
	return getMessageList(socket, command, actionID,
		"MeetmeEntry", "MeetmeListComplete")
}

//	MeetmeMute
//		Mute a Meetme user.
//
func MeetmeMute(socket *Socket, actionID, meetme, usernum string) (map[string]string, error) {
	if len(actionID) == 0 || len(meetme) == 0 || len(usernum) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: MeetmeMute",
		"\r\nActionID: ",
		actionID,
		"\r\nMeetme: ",
		meetme,
		"\r\nUsernum: ",
		usernum,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	MeetmeUnMute
//		Unmute a Meetme user.
//
func MeetmeUnMute(socket *Socket, actionID, meetme, usernum string) (map[string]string, error) {
	if len(actionID) == 0 || len(meetme) == 0 || len(usernum) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: MeetmeUnMute",
		"\r\nActionID: ",
		actionID,
		"\r\nMeetme: ",
		meetme,
		"\r\nUsernum: ",
		usernum,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
