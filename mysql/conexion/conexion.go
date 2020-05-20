package conexion
//import para realizar la conexion
import(
  "database/sql" // Interactuar con bases de datos
	"fmt"   // Liberia que nos permite mostrar datos en la consola
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)
//constantes de la conexion
const (
  //Variables de conexion
  vUser = "root" //user<
  vPass = ""     //password
  vHost = "tcp(127.0.0.1:3306)" //host
  vDataBaseName = "bdalumnos" //nombre de la bd
)

func ConexionBD() (vCon *sql.DB, e error)  {
  //Abrimos la conexion a la base de datos el cual nos regresa la conexion o un error.
  //"vHost=%s vUser=%s vPass=%s vDataBaseName=%s
  vCon, vErr := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s",vUser,vPass,vHost,vDataBaseName))
  if vErr != nil {
    return nil, vErr
  }else{
    return vCon, nil
  }
}
