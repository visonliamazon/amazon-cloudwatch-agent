import { Helmet } from 'react-helmet';
import styled from 'styled-components';
import { Routes ,Route } from 'react-router-dom';

import HomePage from '../HomePage/Loadable';
import Header from '../../components/Header';
import Footer from '../../components/Footer';

const AppWrapper = styled.div`
  max-width: calc(768px + 16px * 2);
  margin: 0 auto;
  display: flex;
  min-height: 100%;
  padding: 0 16px;
  flex-direction: column;
`;


export default function App()  {
    return (
      <AppWrapper>
        <Helmet
          titleTemplate="Amazon CloudWatch Agent"
          defaultTitle="Amazon CloudWatch Agent"
        >
            <meta name="description" content="Amazon CloudWatch Agent" />
        </Helmet>
        <Header />
        <Routes>
            <Route exact path="/" component={HomePage} />
        </Routes>
        <Footer />
      </AppWrapper>
      
)}
