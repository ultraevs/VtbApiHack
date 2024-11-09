import React, { ReactNode } from 'react';
import { NavBar } from "../navbar";
import styles from './styles.module.scss';
import { Header } from '../header';

const AutoLayout: React.FC<{ children: ReactNode }> = ({ children }) => {
  return (
    <div className={styles.autoLayout}>
      <NavBar />
      <div>
        <Header />
        {children}
      </div>

    </div>
  );
};

export { AutoLayout };