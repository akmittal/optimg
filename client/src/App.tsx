import React from 'react';


import './App.css';

import { BrowserRouter as Router } from 'react-router-dom';
import { Layout } from 'antd';

import Routes from './Routes';
import Navbar from './components/Navbar';
import Drawer from './components/Drawer';
import { Content } from 'antd/lib/layout/layout';
import { QueryClient, QueryClientProvider } from 'react-query';


import 'antd/dist/antd.css';
import './index.css';


const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <Layout>
          <Layout>
            <Drawer />
            <Layout className="ml-56 h-100vh">
              <Navbar
                title="Optimg"
                routes={[
                  { breadcrumbName: 'Home', path: '/' },
                  { breadcrumbName: 'Home', path: '/' },
                ]}
              />

              <Content
                className="site-layout-background"
                style={{
                  padding: 24,
                  margin: 0,
                  minHeight: 280,
                }}
              >
                <Routes />
              </Content>
            </Layout>
          </Layout>
        </Layout>
        {/* A <Switch> looks through its children <Route>s and
          renders the first one that matches the current URL. */}
      </Router>
    </QueryClientProvider>
  );
}

export default App;
