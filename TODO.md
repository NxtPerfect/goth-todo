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
- [ ] tasks part 1
    - [x] add new task
        - We do a lot of calls to verify in each function
        should move it to internal file
        and call it when needed
        - generate auth token once
        then save to users table
        - however it returns all of the new html
        while it should only return the new element
        - on "add new task" button, we always return
        new task form
- [ ] On forms, add error messages
    - [ ] Create types with error messages
        in templ check if they're ""
        if no then show error message

        when i change target to error message
        then they show up
        but then it also swaps that element
        for home page on success
- [ ] tasks part 2
    - [ ] editing task
    - [ ] completing task
    - [ ] search task by name/descr

# Frontend
- [ ] home page
    - [ ] number of tasks
    - [ ] top 5 recent tasks
    - [ ] sort by
        - [ ] most progressed tasks
        - [ ] most recent
        - [ ] oldest
        - [ ] least progresed
        - [ ] alphabetical
    - [ ] Task
        - [ ] id
        - [ ] title
        - [ ] description
        - [ ] due date
        - [ ] created date
        - [ ] last modified
- [x] Login/Register page
    - [x] Login/Register form
    - [ ] tailwind styling
- [ ] Profile page
    - [ ] Total tasks finished
    - [ ] total tasks started
        - [ ] show as nice circle graph


# Optimizations
- [ ] switch to fasthttp
- [ ] use echo/air
