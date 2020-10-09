# Data Model

## Database Choice

- This repository uses MongoDB as its database
- An SQL database like PostgreSQL even though CA in nature, is not used because of the unstructured nature of the data makes it difficult to work with an SQL database
- MongoDB is a CP database, with a lot of flexibility provided to the structure of the data being stored which gives us a good balance of ease of usability and availability, consistency and partition tolerance

## Collections

### Users

| Attribute   | Type     | Description                     |
| ----------- | -------- | ------------------------------- |
| uid         | String  | A unique ID for each user       |
| uname       | String   | Username                        |
| pass        | String   | User password                   |
| dateJoined  | Date     | User's join date                |

### SessionData

| Attribute   | Type      | Description                                |
| ----------- | --------- | ------------------------------------------ |
| sessid      | String   | A unique ID for each created session       | 
| lang        | String    | Current active language of the session     |
| private     | Boolean   | Flag for whether session is private        |
| accessIds   | Integer[] | access IDs who have access to this session |
| sessData    | String    | Content of the session                     |

### ActiveSessions

| Attribute   | Type                           | Description                                   |
| ----------- | ------------------------------ | --------------------------------------------- |
| sessid      | String                         | A unique ID for each created session          |
| activeUsers | {String, {Integer, Integer}}[] | Active uid and corresponding cursor positions |

### Messages

| Attribute   | Type      | Description                                |
| ----------- | --------- | ------------------------------------------ |
| msgid       | String    | Unique ID for each message                 |
| sessid      | String    | ID of the owning session                   | 
| msgTime     | datetime  | Time message was sent                      |
| senderId    | String    | uid of sender                              |
| accessIds   | Integer[] | access IDs who have access to this session |
| msgData     | String    | Message's content                          |

### SessIdStore

| Attribute   | Type     | Description                         |
| ----------- | -------- | ----------------------------------- |
| sessid      | String   | A unique ID for each session        |
| used        | Boolean  | Flag for whether the sessId is used |

### UserIdStore

| Attribute   | Type     | Description                         |
| ----------- | -------- | ----------------------------------- |
| uid         | String   | A unique ID for each user           |  
| used        | Boolean  | Flag for whether the sessId is used |
