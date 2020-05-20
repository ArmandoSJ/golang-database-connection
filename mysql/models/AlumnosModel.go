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
