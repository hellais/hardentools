import React from 'react'
import { NavLink, Route } from 'react-router-dom'
import Scan from './Scan'
import Overview from './Overview'

import '../assets/css/main.css'
import HardenToolsLogo from '../assets/images/icon@512.png'

export const colors = {
  darkGrey: "#4A4A4A",
  lightGrey: "#9B9B9B",
  lighterGrey: "#D8D8D8",
  white: "#ffffff",
}

class App extends React.Component {
  render() {
    return (
      <div>
        <div style={{backgroundColor: colors.darkGrey, height: '10vh', display: 'flex'}}>
          <div style={{display: 'flex', alignItems: 'center', paddingLeft: '20px'}}>
            <img src={HardenToolsLogo} height='50px'/>
            <div style={{display: 'flex', flexDirection: 'column', paddingLeft: '20px' }}>
              <h2 style={{color: colors.white, margin: '0'}}>Harden Tools</h2>
              <span>v1.1.1</span>
            </div>
          </div>
        </div>

        <div style={{display: 'flex', minHeight: '90vh'}}>
          <div style={{backgroundColor: colors.lightGrey, width: '183px'}} className='menu-items'>
            <NavLink to="/overview" activeClassName="active">
              <div className='menu-item-container'>
                <div style={{margin: '0 auto', width: '77px', height: '77px', backgroundColor: 'black', borderRadius: '70px'}} />
                <h2>Overview</h2>
              </div>
            </NavLink>
            <NavLink to="/scan" activeClassName="active">
              <div className='menu-item-container'>
                <div style={{margin: '0 auto', width: '77px', height: '77px', backgroundColor: 'black', borderRadius: '70px'}} />
                <h2>Scan</h2>
              </div>
            </NavLink>
          </div>

          <div style={{backgroundColor: colors.white}}>
            <div>
              <Route path="/scan" component={Scan}/>
              <Route path="/overview" component={Overview}/>
            </div>
          </div>
       </div>
      </div>
    );
  }
}

export default App;
