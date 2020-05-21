package conexion
//import para realizar la conexion
import(
  "database/sql" // Interactuar con bases de datos
	_ "github.com/lib/pq" // La librería que nos permite conectar a postgresql
  "fmt"
)
//constantes de la conexion
const (
  //Variables de conexion
 host = "localhost"
 user = "postgres"
 password = "hola"
 databaseName ="postgres"


)

func ConexionBD() (vCon *sql.DB, e error)  {
  //Abrimos la conexion a la base de datos el cual nos regresa la conexion o un error.
  //"vHost=%s vUser=%s vPass=%s vDataBaseName=%s
  //vConnectionString := "postgres://"+vUser+":"+vPwd+"@"+vHost+":"+vPort+"/"+vDataBaseName+"?sslmode=disable"
  vConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)
  vCon, vErr := sql.Open("postgres", vConnectionString)
  if vErr != nil {
    return nil, vErr
    }else{
      return vCon, nil
    }

}
