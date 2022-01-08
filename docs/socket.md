# Socket Commands

## Backend Commands

[//]: # (handlers_start)
*Generated on 08.01.2022 22:51:29*

### CHANGE_PREF

> Changes a game-setting

`CHANGE_PREF { "key": (key!), "value": (value!)" }`  
👉 args: [1 <= x]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### CHAT

> Sends a chat message

`CHAT (...message!)`  
👉 args: [1 <= x]
!!! danger "Access"
	- [ ] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### EXPLAIN 🔰[^1]

> Returns help for a handler

`EXPLAIN (handler!)`  
👉 args: [x == 1]
!!! danger "Access"
	- [x] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### JOIN

> Joins a game with a username

`JOIN (username!)`  
👉 args: [x == 1]
!!! danger "Access"
	- [x] Guest
	- [ ] Joined
	- [ ] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### LIST

> Returns a list of all connected players from the current game

`LIST`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### NEXT_ROUND

> Starts the next round (displays the next topic)

`NEXT_ROUND`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [ ] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [x] Show Vote Results

---

### SKIP

> Skips to the next game circle

`SKIP`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [ ] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### START

> Starts a game

`START`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### STATS

> Returns the points for each player

`STATS`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [ ] Lobby
	- [x] Submit GIF
	- [ ] Cast Votes
	- [x] Show Vote Results

---

### SUBMIT_GIF

> Submits a GIF

`SUBMIT_GIF (url!)`  
👉 args: [x == 1]
!!! danger "Access"
	- [ ] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [ ] Lobby
	- [x] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### TOPIC_ADD

> Adds a new topic to the game

`TOPIC_ADD (topic!)`  
👉 args: [1 <= x]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### TOPIC_ADD_ALL

> Adds all topics from a JSON array

`TOPIC_ADD_ALL (...topics: Array<string>!)`  
👉 args: [1 <= x]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### TOPIC_CLEAR

> Removes all topics from the current game

`TOPIC_CLEAR`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### TOPIC_LIST

> Returns a list with all topics from the current game

`TOPIC_LIST`  
👉 args: [x == 0]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

---

### TOPIC_REMOVE

> Removes a topic from the game

`TOPIC_REMOVE (topic!)`  
👉 args: [1 <= x]
!!! danger "Access"
	- [ ] Guest
	- [ ] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [ ] Submit GIF
	- [ ] Cast Votes
	- [ ] Show Vote Results

---

### VOTE

> Votes for a submission

`VOTE (url!)`  
👉 args: [x == 1]
!!! danger "Access"
	- [ ] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [ ] Lobby
	- [ ] Submit GIF
	- [x] Cast Votes
	- [ ] Show Vote Results

---

### WHOAMI 🔰[^1]

> Returns information about the client and game

`WHOAMI`  
👉 args: [x == 0]
!!! danger "Access"
	- [x] Guest
	- [x] Joined
	- [x] Leader

!!! hint "States"
	- [x] Lobby
	- [x] Submit GIF
	- [x] Cast Votes
	- [x] Show Vote Results

[//]: # (handlers_end)

[^1]: Dev-Only

## Frontend Commands

ABC
