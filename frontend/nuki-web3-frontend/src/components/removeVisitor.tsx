import React, { useState } from 'react';
import { removeVisitor } from '../functions/functions';

type Props = {
  handleMasterLogsChange: (data: any) => void;
};

type State = {
  [key: string]: string
};




export default function RemoveVisitor(props : Props) {
  const { handleMasterLogsChange } = props
  const [state, setState] = useState<State>({});
    
  const handleStateChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setState((prevState) => {
      return {
        ...prevState,
        [name]: value
      }
    })
    
  }
    
  
  const handleLogsClick = async (event: any) => {
    const data = await removeVisitor(state?.visitor)
    const outcome = data === "removeVisitor failed" ? "failed" : `success: ${JSON.stringify(data)}`
    handleMasterLogsChange(["removeVisitor", outcome])
  }

  return (
    <div className="function-box">
      <div className="box-heading">
        <h1>Remove Visitor</h1>
        <span className="text-extra"><p>removeVisitor</p></span><p>Function inputs:</p>
          <p>(<span className="text-extra">address visitor:</span> string)</p>
      </div>
          
      <div className="box-inputs">
        <p>visitor</p>
        <input aria-label="visitor" name="visitor" onChange={handleStateChange} type="string" placeholder="visitor"/>
      </div>

        <button className="box-button" onClick={handleLogsClick}>removeVisitor</button>
    </div>
  )
}
