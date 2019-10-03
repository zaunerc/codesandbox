/** @type WebpackConfig[] */
const configs = [
	{
		entry: {
			browser: './src/client.tsx',
		},
		output: {
			path: __dirname + '/dist',
			filename: '[name].js',
		},
		// Currently we need to add '.ts' to the resolve.extensions array.
		resolve: {
			extensions: ['.ts', '.tsx', '.js', '.jsx'],
		},

		// Source maps support ('inline-source-map' also works)
		devtool: 'source-map',

		module: {
			rules: [
				{
					test: /\.tsx?$/,
					loader: 'ts-loader',
				},
			],
		},
	}
];

module.exports = configs;
