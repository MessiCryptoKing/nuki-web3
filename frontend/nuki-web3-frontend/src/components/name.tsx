import React from 'react';
import { name } from '../functions/functions';

type Props = {
  handleMasterLogsChange: (data: any) => void;
};




export default function Name(props : Props) {
  const { handleMasterLogsChange } = props
  
  
  const handleLogsClick = async (event: any) => {
    const data = await name()
    const outcome = data === "name failed" ? "failed" : `success: ${JSON.stringify(data)}`
    handleMasterLogsChange(["name", outcome])
  }

  return (
    <div className="function-box">
      <div className="box-heading">
        <h1>Name</h1>
        <span className="text-extra"><p>name</p></span>
      </div>
          
        <button className="box-button" onClick={handleLogsClick}>name</button>
    </div>
  )
}
