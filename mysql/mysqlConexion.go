package main
//Conexion a mysql
import(
    "fmt"
    "./conexion"
    "./models"
)

func main(){
  getAlumnos()
}

func getAlumnos()  {
  vCon, vErr := conexion.ConexionBD()
    if vErr != nil {
      fmt.Printf("No se ha podido encontrar la base de datos", vErr)
    }else{
      alumnosModel := models.AlumnosModel{
         Db: vCon,
      }
      alumnos, err := alumnosModel.FindAll()
      if err != nil{
        fmt.Println("Error al encontrar alumnos", err)
      }else{
        fmt.Println(alumnos)
        fmt.Println("Lista de alumnos")
        for _, alumno := range alumnos{
          fmt.Println("Numero de control",alumno.NoControl)
          fmt.Println("Carrera",alumno.Carrera)
          fmt.Println("Nombre",alumno.Nombre)
          fmt.Println("Telefono",alumno.Telefono)
        }
      }
    }
}
