package main
import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "digital_twin/db"
    "digital_twin/token"
)
const mongodb_url = "mongodb://localhost:27017/"
func env_test(){
    url := "mongodb://localhost:27017/"
    db_connection , err := db.Mongo_Connect(url)
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println("Connected")
    }
    databases ,err := db.List_Databases(db_connection)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Databases:", databases)
    // collection
    collections, err := db.List_Collections("user-data", db_connection)
    if err != nil {
            fmt.Println(err)
            return
    }
    fmt.Println("Collections:", collections)
    fmt.Println()
    fmt.Println("Fetching the data from database ")
    obj_id := "67499573de39ae54d7e9496a"
    database := "user-data"
    collection := "records"
    documents , err := db.Get_Document_by_id(obj_id,database, collection, db_connection)
    if err != nil {
        fmt.Println(err)
        return
    }else{
        fmt.Println(documents)
    }
    name := "Kavi"
    password := "mcw@123"
    role := "read"
    db := "digital_twin"
    gen_token , err := token.Generate_Token(name, password, role, db)
    if err != nil{
            fmt.Print(err)
    }else {
            fmt.Println("Generated Token.")
            fmt.Println(gen_token)
    }
    fmt.Println("Validting the Token")
    rev_token, err := token.Validate_Token(gen_token)
    if err != nil{
            fmt.Println(err)
    }else{
            fmt.Println("Username:", rev_token.Username)
            fmt.Println("Role:", rev_token.Role)
            fmt.Println("Database:", rev_token.DB)
    }
}
/* {
  acknowledged: true,
  insertedId: ObjectId('67499573de39ae54d7e9496a')
}
*/
// api part starts
type login_data struct{
    Username string  `json:"username"`
    Password string `json:"password" `
}
func login(c *gin.Context){
    var user_data login_data
    if err := c.ShouldBindJSON(&user_data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // creating the connection to the mongodb db datbase
    db_con , err := db.Mongo_Connect(mongodb_url)
    if err != nil{
        fmt.Println(err)
    }
    // geting the data from datbase using name
    user_data_db, err  := db.Get_data_by_username(user_data.Username, db_con)
    if err != nil{
        fmt.Println("There is an ERROR: ",err)
        // have to create an return statement
    }
	fmt.Println(user_data_db)
    // checking the password
    
    // if the user name and passwords matchs our database then have to generate the token and return it
    // if the password doesnt math with user name handle it
    // if the user is not in the data handle it
    c.JSON(http.StatusOK, gin.H{
        "message":  "Login successful",
        "username": user_data.Username,
    })
}
func main(){
    env_test()
    router := gin.Default()
    router.POST("/login",login)
    router.Run(":8080")
}
