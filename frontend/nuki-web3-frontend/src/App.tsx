
import React, { useState } from 'react';
import './App.css';
import Footer from './Footer';
import Details from './Details';
import changeTheme from './changeTheme';
import Nav from './Nav';
import AddVisitor from './components/addVisitor';
import IsSenderAllowedToEnter from './components/isSenderAllowedToEnter';
import RemoveVisitor from './components/removeVisitor';
import Name from './components/name';

type MasterLogs = any[];

export default function App() {
  const [theme, setTheme] = useState("dark")
  const [masterLogs, setMasterLogs] = useState<MasterLogs>([]);

  const handleMasterLogsChange = (data: any) => {
    setMasterLogs(prevState => [data, ...prevState])
  }

  const handleTheme = () => {
    setTheme(changeTheme(theme))
  }
  return (
    <div className="App">
    <Nav handleTheme={handleTheme} />
    <Details masterLogs={masterLogs} />
      <div className="components">
              <AddVisitor handleMasterLogsChange={handleMasterLogsChange} />
      <IsSenderAllowedToEnter handleMasterLogsChange={handleMasterLogsChange} />
      <RemoveVisitor handleMasterLogsChange={handleMasterLogsChange} />
      <Name handleMasterLogsChange={handleMasterLogsChange} />
      </div>
    <Footer />
    </div>
  )
}
