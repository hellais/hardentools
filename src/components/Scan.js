import React from 'react'

export default class Scan extends React.Component {
  constructor(props) {
    super(props)
    this.state =  {
      systemName: ''
    }
    this.startScan = this.startScan.bind(this)
  }

  startScan() {
    var execSync = require('child_process').execSync
    var result = execSync('uname -a')
    console.log(result)
    this.setState({systemName: result})
  }

  render() {
    const {
      systemName
    } = this.state
    return (
      <div>
      XXXX
        <div className="progress-bar" style={{backgroundColor: 'black'}}>
          <div style={{backgroundColor: 'red', width: '100px'}}>X</div>
        </div>
        <p>{systemName}</p>
        <button onClick={() => this.startScan()}>Do it</button>
      </div>
    )
  }
}
