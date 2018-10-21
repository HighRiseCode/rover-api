import React from 'react';

class Box extends React.Component{
  render(){
    return(
      <div className="wrapper">
        <div className="box">
          <img src={this.props.color} alt=""/>
        </div>
        <button onClick={()=>{this.props.handleClick()}}>Change Color</button>
      </div>
    )
  }
}

export default Box;
