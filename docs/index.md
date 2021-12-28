# General Concepts

## Communication

**GAME:**

| Field           | Description                                      |
| --------------- | ------------------------------------------------ |
| GameID          |                                                  |
| Clients         |                                                  |
| LastInteraction | used for janitor service to clean up old session |

**PAP:**

* client: create new game
    * client: POST request
        * server: generate unique ID
* client: change url with generated game id (to share)

* IF: client: connected per socket
    * server: if connected client was first, make 'em leader
    * server: add client to game
    * server: send `CLIENT_CONNECT <username>` to all clients

## Socket Messages

### Client To Server

#### C_JOIN <username>
Server should send [S_PLAYER_JOINED](#s_player_joined-username) to all clients

#### C_LEAVE
Server should send [S_PLAYER_LEFT](#s_player_left-username) to all clients

#### C_PLAYER_ROUND_SUBMIT_GIF <URL>
Server should send [S_ROUND_PLAYER_SUBMITTED_GIF](#s_round_player_submitted_gif-username) to all clients

#### C_PLAYER_ROUND_SUBMIT_VOTE <vote-id>
Server should send [S_ROUND_PLAYER_SUBMITTED_VOTE](#s_round_player_submitted_vote-username) to all clients

#### C_LEADER_ROUND_NEXT
Server should send [S_ROUND_LEADER_NEXT](#s_round_leader_next) to all clients

---

### Server To Client

#### S_PLAYER_JOINED <username>
Clients should display the joined player

#### S_PLAYER_LEFT <username>
Clients should abandon the left player

#### S_ROUND_PLAYER_SUBMITTED_GIF <username>
Clients should remove player from "waiting for gif" list

#### S_ROUND_PLAYER_SUBMITTED_VOTE <username>
Clients should remove player from "waiting for vote" list

#### S_ROUND_NEXT <topic>
Clients should display the next round

## Contributing to the docs
- install mkdocs and additional plugins with `python3 -m pip install -r requirements.txt`
- launch the local development server with `python3 -m mkdocs serve`
