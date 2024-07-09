import React from 'react';

const styles = {
  footer: {
    textAlign: 'center', 
    padding: '20px', 
    marginTop: '20px', 
    backgroundColor: '#f0f0f0',
  },
  ul: {
    listStyleType: 'none', 
    padding: 0,
  },
  li: {
    display: 'inline', 
    marginRight: '10px',
  },
  link: {
    textDecoration: 'none',
  },
};

const Footer = () => {
  return (
    <footer style={styles.footer}>
      <p>ExpenseTracker © 2023</p>
      <p>All Rights Reserved</p>
      <nav>
        <ul style={styles.ul}>
          <li style={styles.li}>
            <a href="#" style={styles.link}>Home</a>
          </li>
          <li style={styles.li}>
            <a href="#" style={styles.link}>About</a>
          </li>
          <li style={styles.li}>
            <a href="#" style={styles.link}>Contact Us</a>
          </li>
        </ul>
      </nav>
    </footer>
  );
};

export default Footer;