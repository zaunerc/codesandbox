export class DataSource {

	private static counter = 0;

	private static sleep = (milliseconds: number) => {
		return new Promise(resolve => setTimeout(resolve, milliseconds));
	}
	  
	public static async fetchData(idToFetch: number): Promise<string> {
		
		console.log(`${new Date().toString()} fetchData called with ${idToFetch} as idToFetch`);

		let sleepTime: number;
		let returnValue: string;

		switch (DataSource.counter) {
			case 0:
				sleepTime = 10000;
				returnValue = `Return value from FIRST promise (idToFetch was ${idToFetch})`;
				break;
		
			case 1:
				sleepTime = 2000;
				returnValue = `Return value from SECOND promise (idToFetch was ${idToFetch})`;
				break;

			default:
				throw new Error("Invalid application state");
		}

		DataSource.counter++;
		
		await DataSource.sleep(sleepTime);
		return returnValue;
	}

}
