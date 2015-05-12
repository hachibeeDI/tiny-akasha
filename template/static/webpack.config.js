module.exports = {
  entry: './temp/entry.js',
  output: {
    path: __dirname + '/dist/',
    filename: 'bundle.js',
    //at this directory our bundle file will be available
    //make sure port 8090 is used when launching webpack-dev-server
    publicPath: 'http://localhost:8080/js/'
  },
  module: {
    loaders: [
      { test: /\.coffee$/, loader: "coffee-loader" },
      { test: /\.ts$/, loader: 'typescript-loader' },
      {
        //tell webpack to use jsx-loader for all *.jsx files
        test: /\.jsx$/,
        loader: 'jsx-loader?insertPragma=React.DOM&harmony'
      }
    ]
  },
  externals: {
    //don't bundle the 'react' npm package with our bundle.js
    //but get it from a global 'React' variable
    'react': 'React',
  },
  resolve: {
    extensions: ['', '.js', '.jsx', '.coffee', '.ts']
  }
};
