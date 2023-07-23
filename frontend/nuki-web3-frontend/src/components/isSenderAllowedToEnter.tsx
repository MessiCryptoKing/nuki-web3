import React, { useState } from 'react';
import { isSenderAllowedToEnter } from '../functions/functions';

type Props = {
  handleMasterLogsChange: (data: any) => void;
};

type State = {
  [key: string]: string
};




export default function IsSenderAllowedToEnter(props : Props) {
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
    const data = await isSenderAllowedToEnter(state?.home)
    const outcome = data === "isSenderAllowedToEnter failed" ? "failed" : `success: ${JSON.stringify(data)}`
    handleMasterLogsChange(["isSenderAllowedToEnter", outcome])
  }

  return (
    <div className="function-box">
      <div className="box-heading">
        <h1>Is Sender Allowed To Enter</h1>
        <span className="text-extra"><p>isSenderAllowedToEnter</p></span><p>Function inputs:</p>
          <p>(<span className="text-extra">address home:</span> string)</p>
      </div>
          
      <div className="box-inputs">
        <p>home</p>
        <input aria-label="home" name="home" onChange={handleStateChange} type="string" placeholder="home"/>
      </div>

        <button className="box-button" onClick={handleLogsClick}>isSenderAllowedToEnter</button>
    </div>
  )
}
