import React from 'react';

class Box extends React.Component{
  render(){
    return(
      <div className="wrapper">
        <div
          style={{backgroundColor:`${this.props.color}`}}
          className="box">
        </div>
        <button onClick={()=>{this.props.handleClick()}}>Change Color</button>
      </div>
    )
  }
}

export default Box;
