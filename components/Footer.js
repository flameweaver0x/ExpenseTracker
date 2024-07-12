import React, 'react';

class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    // Update state so the next render will show the fallback UI
    return { hasError: true };
  }

  componentDidCatch(error, errorInfo) {
    // You can also log the error to an error reporting service
    console.error("Error caught by ErrorBoundary: ", error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      return <h2>Something went wrong.</h2>;
    }

    return this.props.children; 
  }
}

import React from 'react';
import Footer from './Footer';
import ErrorBoundary from './ErrorBoundary'; // Assuming ErrorBoundary is saved in the same directory

function App() {
  return (
    <div className="App">
      <ErrorBoundary>
        <Footer />
        {/* Any other component that you want to protect */}
      </ErrorBoundary>
    </div>
  );
}

export default App;