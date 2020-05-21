package models

import(
  "database/sql"
  "../domain"
)
type AlumnosModel struct{
    Db *sql.DB

}

func (alumnosModel AlumnosModel)  FindAll() ([]domain.Alumnos, error){
  vRows, vErr :=  alumnosModel.Db.Query("Select *from alumno")
  if vErr != nil{
    return nil, vErr
  }else{
    vAlumnos := []domain.Alumnos{}
    for vRows.Next(){
      var NoControl string
      var Carrera string
      var Nombre string
      var Telefono string

      err2 := vRows.Scan(&NoControl,&Carrera,&Nombre,&Telefono)
      if err2 != nil{
        return nil, err2
      }else{
        vAlumno := domain.Alumnos{NoControl,Carrera,Nombre,Telefono}
        vAlumnos = append(vAlumnos,vAlumno)
      }
    }

    return vAlumnos, nil
  }
}

//Encontrar alumno con id
func (alumnosModel AlumnosModel)  FindWithID(vID int) (domain.Alumnos, error){
  vRows, vErr :=  alumnosModel.Db.Query("Select *from alumno where NoControl= ?", vID)
  if vErr != nil{
    return domain.Alumnos{}, vErr
  }else{
    vAlumno := domain.Alumnos{}
    for vRows.Next(){
      var NoControl string
      var Carrera string
      var Nombre string
      var Telefono string

      err2 := vRows.Scan(&NoControl,&Carrera,&Nombre,&Telefono)
      if err2 != nil{
        return domain.Alumnos{}, err2
      }else{
        vAlumno = domain.Alumnos{NoControl,Carrera,Nombre,Telefono}
      }
    }

    return vAlumno, nil
  }
}

//Crear un alumno
func (alumnosModel AlumnosModel)  InsertAlumno(alumno *domain.Alumnos) error{
  vResult, vErr :=  alumnosModel.Db.Exec("insert into alumno(NoControl,Carrera,Nombre,Telefono)values(?,?,?,?)",alumno.NoControl,
  alumno.Carrera,alumno.Nombre,alumno.Telefono)
  if vErr != nil{
    return vErr
  }else{
    vResult.LastInsertId();
    //edad, _ = vResult.LastInsertString();
    return nil
  }
}

//Update un alumno
func (alumnosModel AlumnosModel)  UpdateAlumno(alumno domain.Alumnos) (int64 ,error){
  vResult, vErr :=  alumnosModel.Db.Exec("update alumno set Carrera = ?, Nombre = ?, Telefono = ? where NoControl = ?",
  alumno.Carrera,alumno.Nombre,alumno.Telefono,alumno.NoControl)
  if vErr != nil{
    return 0,vErr
  }else{
    return vResult.RowsAffected()
  }
}

//Update un alumno
func (alumnosModel AlumnosModel)  DeleteAlumno(NoControl string) (int64 ,error){
  vResult, vErr :=  alumnosModel.Db.Exec("delete from alumno where NoControl = ?",NoControl)
  if vErr != nil{
    return 0,vErr
  }else{
    return vResult.RowsAffected()
  }
}
