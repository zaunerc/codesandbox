import React from "react";
import { DataSource } from "./DataSource";

interface DisplayComponentProps {
	idToFetch: number;
}

interface DisplayComponentState {
	returnValueFromPromise: string;
}

export class DisplayComponent extends React.Component<DisplayComponentProps, DisplayComponentState> {
	
	constructor(props: DisplayComponentProps) {
		super(props);

		this.state = {
			returnValueFromPromise: "No promise has returned a value yet",
		};
	}

	async componentDidMount() {
		const result = await DataSource.fetchData(this.props.idToFetch);
		this.setState({
			returnValueFromPromise: result
		});
	}
	
	async componentDidUpdate(prevProps: DisplayComponentProps) {
		console.debug(`${new Date().toString()} componentDidUpdate called (prevProps: ${JSON.stringify(prevProps)}, currentProps: ${JSON.stringify(this.props)})`);
		if (prevProps.idToFetch != this.props.idToFetch) {
			const result = await DataSource.fetchData(this.props.idToFetch);
			this.setState({
				returnValueFromPromise: result
			});
		}
	}

	render() {
		return (
			<ul>
				<li>DisplayComponent: Latest value received for prop idToFetch ===> {this.props.idToFetch}</li>
				<li>DisplayComponent: Return value from promise ===> <span style={{color: "green"}}>{this.state.returnValueFromPromise}</span></li>
			</ul>
		);
	}
}
