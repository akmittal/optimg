import React, { ReactElement } from 'react'
import { Route, Switch } from 'react-router-dom'
import Home from './components/Home';
import Navbar from './components/About';
import Optimize from './pages/Optimize';
import Gallery from './pages/Gallery';
import Compare from './pages/Compare';
import Image  from './pages/Image';

interface Props {
    
}

function Routes({}: Props): ReactElement {
    return (
        <Switch>
        <Route path="/about">
         
        </Route>
        <Route path="/optimize">
          <Optimize />
        </Route>
        <Route path="/gallery/:path?/:pageNo?">
          <Gallery />
        </Route>
        <Route path="/image/:path?/:name?">
          <Image />
        </Route>
        <Route path="/compare">
          <Compare />
        </Route>

        <Route path="/">
          <Home />
        </Route>
      </Switch>
    )
}

export default Routes
