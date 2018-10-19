import React from 'react';
import Box from '../component/box.js';
import {connect} from "react-redux";
import * as actionCreators from "../actions/index.js";

class BoxCon extends React.Component{
  render(){
    return(
      <Box handleClick={this.props.loadColor} color={this.props.color}></Box>
    )
  }
}

const mapStateToProps=(state)=>{
  return state
};

export default connect (mapStateToProps, actionCreators)(BoxCon)
