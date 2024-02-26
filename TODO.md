# Backend
- [x] connect to database
- [x] login user
    - [x] search db for username + password
    - [x] come up with auth_token generation
- [x] get tasks from database
    - [x] currently errors
    - [x] perhaps the "insert" file didn't get run?
        there are no databases created
    - [x] return tasks for logged in user
- [x] signup user
- [/] tasks part 1
    - [x] add new task
        - [x] on "add new task" button, we always return
        new task form
        - [ ] We do a lot of calls to verify in each function
        should move it to internal file
        and call it when needed
        generate auth token once
        then save to users table
        - however it returns all of the new html
        while it should only return the new element
- [x] On forms, add error messages
- [/] tasks part 2
    - [x] editing task
        - [ ] ability to cancel editing
    - [ ] completing task
    - [ ] search task by name/descr

# Frontend
- [/] home page
    - [x] number of tasks
    - [ ] top 5 recent tasks
    - [ ] sort by
        - [ ] most progressed tasks
        - [ ] most recent
        - [ ] oldest
        - [ ] least progresed
        - [ ] alphabetical
    - [/] Task
        - [x] title
        - [x] description
        - [ ] due date
        - [ ] created date
        - [ ] last modified
- [x] Login/Register page
    - [x] Login/Register form
    - [x] tailwind styling
- [ ] Profile page
    - [ ] Total tasks finished
    - [ ] total tasks started
        - [ ] show as nice circle graph
- [ ] loading cursor/spinner svg


# Optimizations
- [ ] switch to fasthttp
- [x] use air
- [ ] use echo?
