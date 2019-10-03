import React from "react";
import { DisplayComponent } from "./DisplayComponent";

interface AppState {
	idToFetch: number;
	appLog: string[]
}

export class App extends React.Component<{}, AppState> {
	
	private initialIdToFetch = 1;
	private initalAppLog = [];
	
	private nextIdToFetch = 2;

	constructor(props: {}) {
		super(props);

		this.state = {
			idToFetch: this.initialIdToFetch,
			appLog: this.initalAppLog,
		}
	}
	
	componentDidMount() {
		this.scheduleLogMessages();
		setTimeout( () => {
			this.setState({
				idToFetch: this.nextIdToFetch
			});	
		}, 3000);
	}

	render() {
		return (
			<>
				<h1>SUMMARY</h1>
				<p>
					This proof of concept shows that it is important to think about cancelling
					promises as soon as they are not needed anymore. Otherwise this might
					lead to an inconsistent UI state.

					Though the second promise is started after the first promise the result
					the DisplayComponent is going to show in the end will be the result from the
					first promise.
					
					This proof of concept takes about 10 seconds to complete. To restart the
					PoC just reload the site.
				</p>
				<h1>DISPLAY COMPONENT</h1>
				<DisplayComponent idToFetch={this.state.idToFetch}/>
				<h1>EXPLANATORY MESSAGES TO SEE WHAT'S GOING ON</h1>
				<ul>
					{this.state.appLog.map( (msg) => <li key={msg}>{msg}</li> )}
				</ul>
			</>
		);
	}

	private scheduleLogMessages() {
		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} STARTED PROOF OF CONCEPT.`)});	
		}, 0);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Loaded DisplayComponent with initial props (idToFetch = ${this.initialIdToFetch}).`)});	
		}, 0);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Started first promise.`)});	
		}, 0);
		
		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Sent new props to DisplayComponent (idToFetch = ${this.nextIdToFetch}).`)});	
		}, 3000);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Started second promise.`)});	
		}, 3000);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} First promise has been resolved.`)});	
		}, 6000);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Second promise has been resolved.`)});	
		}, 10000);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} Now the DisplayComponent displays the wrong return value.`)});	
		}, 10000);

		setTimeout( () => {
			this.setState({appLog: this.state.appLog.concat(`${new Date().toString()} FINISHED PROOF OF CONCEPT.`)});	
		}, 10000);
	}
}
