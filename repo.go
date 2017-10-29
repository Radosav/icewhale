package main

var currentId int


// Give us some seed data
func init() {
    UserCreate(User{
        Name: "Radosav",
        Desc: "Totally Awsome Coder",
        Status: true,
        Cost:"1 coffee/h, 2 Packs of cigarettes/day",
        Mail:"rasabrajic@gmail.com",
        Role:"Admin",
        Position:"Api Coder",
        Image:"https://avatars1.githubusercontent.com/u/6116078?s=400&u=85e96b04d88f1dddb7b5603bdcbab1002cadc71c&v=4"})


    UserCreate(User{
        Name: "Ahmet",
        Desc: "Pro. Sys admin and Security Expert ( Hacker )",
        Status: true,
        Cost:"0.75 coffee/h, 1.5 Packs of cigarettes/day",
        Mail:"ahmet.gudenoglu@gmail.com",
        Role:"Admin",
        Position:"System Admin",
        Image:"http://3.images.southparkstudios.com/blogs/southparkstudios.com/files/2014/02/1008-MWBZ-faq-q1.jpg?quality=0.8"})


    UserCreate(User{
        Name: "Krle",
        Desc: "Dev in traning",
        Status: true,
        Cost:"Free",
        Mail:"krstickrl3@gmail.com",
        Role:"Developer",
        Position:"Frontend Developer",
        Image:"https://www.biography.com/.image/t_share/MTM2OTI2NTY2Mjg5NTE2MTI5/justin_bieber_2015_photo_courtesy_dfree_shutterstock_348418241_croppedjpg.jpg"})

    TaskCreate(Task{Name: "Create Design", Assignee:UserFind("Krle")})
    TaskCreate(Task{Name: "Set up Server", Assignee:UserFind("Ahmet")})
    TaskCreate(Task{Name: "Write API", Assignee:UserFind("Radosav")})
}
