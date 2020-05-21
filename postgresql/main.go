package main

import(
  "fmt"
  "./conexion"
  "./models"
  "./domain"
)
func main(){
  GetAlumnos()
  //GetAlumno()
  //InsertAlumno()
  //UpdateAlumno()
  //DeleteAlumno()
}



func DeleteAlumno()  {
  vCon, vErr := conexion.ConexionBD()
    if vErr != nil {
      fmt.Printf("No se ha podido encontrar la base de datos", vErr)
    }else{
      alumnosModel := models.AlumnosModel{
         Db: vCon,
      }
      NoControl := "4444"
      vRows, err := alumnosModel.DeleteAlumno(NoControl)
      if err != nil{
        fmt.Println(err)
      }else{
        if vRows > 0{
          fmt.Println("Se ha eliminado correctamente el numero de control :", NoControl)
        }
      }

    }
}

func UpdateAlumno()  {
  vCon, vErr := conexion.ConexionBD()
    if vErr != nil {
      fmt.Printf("No se ha podido encontrar la base de datos", vErr)
    }else{
      alumnosModel := models.AlumnosModel{
         Db: vCon,
      }
      alumno := domain.Alumnos{
        NoControl: "4444",
        Carrera: "Tics2",
        Nombre: "Taylor",
        Telefono: "5812256",

      }
      vRows, err := alumnosModel.UpdateAlumno(alumno)
      if err != nil{
        fmt.Println(err)
      }else{
        if vRows > 0{
          fmt.Println("Se ha actualizado correctamente el numero de control :", alumno.NoControl)
        }
      }

    }
}

func InsertAlumno()  {
  vCon, vErr := conexion.ConexionBD()
    if vErr != nil {
      fmt.Printf("No se ha podido encontrar la base de datos", vErr)
    }else{
      alumnosModel := models.AlumnosModel{
         Db: vCon,
      }
      alumno := domain.Alumnos{
        NoControl: "4444",
        Carrera: "Tics",
        Nombre: "asalazarj",
        Telefono: "548544741",
      }
      err := alumnosModel.InsertAlumno(&alumno)
      if err != nil{
        fmt.Println(err)
      }else{
        fmt.Println("Se ha insertado correctamente el alumno con numero de control:", alumno.NoControl)
      }

    }
}

func GetAlumnos()  {
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


func GetAlumno()  {
  vCon, vErr := conexion.ConexionBD()
    if vErr != nil {
      fmt.Printf("No se ha podido encontrar la base de datos", vErr)
    }else{
      alumnosModel := models.AlumnosModel{
         Db: vCon,
      }
      alumno, err := alumnosModel.FindWithID(15980598)
      if err != nil{
        fmt.Println("Error al encontrar alumnos", err)
      }else{
        fmt.Println(alumno)
        fmt.Println("Lista de alumnos")
        fmt.Println("Numero de control",alumno.NoControl)
        fmt.Println("Carrera",alumno.Carrera)
        fmt.Println("Nombre",alumno.Nombre)
        fmt.Println("Telefono",alumno.Telefono)

      }
    }
}
