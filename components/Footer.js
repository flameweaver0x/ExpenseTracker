import React from 'react';

const Footer = () => {
  return (
    <footer style={{ textAlign: 'center', padding: '20px', marginTop: '20px', backgroundColor: '#f0f0f0' }}>
      <p>ExpenseTracker © 2023</p>
      <p>All Rights Reserved</p>
      <nav>
        <ul style={{ listStyleType: 'none', padding: 0 }}>
          <li style={{ display: 'inline', marginRight: '10px' }}>
            <a href="#" style={{ textDecoration: 'none' }}>Home</a>
          </li>
          <li style={{ display: 'inline', marginRight: '10px' }}>
            <a href="#" style={{ textDecoration: 'none' }}>About</a>
          </li>
          <li style={{ display: 'inline', marginRight: '10px' }}>
            <a href="#" style={{ textDecoration: 'none' }}>Contact Us</a>
          </li>
        </ul>
      </nav>
    </footer>
  );
};

export default Footer;