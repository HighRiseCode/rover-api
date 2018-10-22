import axios from "axios";

export function loadColor(){
  return(dispatch)=>{
    return axios.get("/rover").then((response)=>{
      dispatch(changeColor(response.data.Img));
    })
  }
}

export function changeColor(color){
  return{
    type: "CHANGE_COLOR",
    color: color
  }
}
