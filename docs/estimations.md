# Capacity Estimations

This document will cover some rough estimates of the traffic, storage and bandwidth.

## Traffic

### Assumptions

- 500 users/year, 100 daily active
- 200 sessions/day
- 40 messages/session
- 5000 characters on average typed in a session in a day
- Reads: Writes for sessions and messages = 2:1

### Calculations

- Editor write traffic = (200 * 5000)/(3600 * 24) ~ 10 writes/sec
- Editor Read traffic = 2 * 10 = 20 reads/sec
- Message write traffic = 40 * 200 / (3600 * 24) ~ 0.1 writes/sec
- Message read traffic = 2 * 0.1 = 0.2 reads/sec

## Storage

### Assumptions

- Keep sessions valid for 1 month
- User data kept over 5 years
- Messages kept for as long as the session is active
- An average paste has about 40 messages
- User data is of an average size of 500 bytes
- A message is of an average size of 200 bytes

### Calculations

- Session Data = 5000 bytes * 200 * 30 = 30 MB
- User information = 500 * 500 bytes * 5 = 2 MB
- Messages = 200 bytes * 40 * 200 * 30 = 48 MB

## Bandwidth

### Calculations

- Editor writes = 10 /s * 1 byte = 10 bytes/sec
- Editor reads = 20 bytes/sec
- Message writes = 200 bytes * 0.1 writes/sec = 20 bytes/sec
- Message reads = 40 bytes/sec