import React, { ReactElement } from "react";
import Home from "./Home";
import About from "./About";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Optimize from "./Optimize";

import {
  TopAppBar,
  TopAppBarRow,
  TopAppBarTitle,
  TopAppBarSection,
  TopAppBarFixedAdjust,
} from "@rmwc/top-app-bar";
import {
  Drawer,
  DrawerTitle,
  DrawerContent,
  DrawerAppContent,
  DrawerHeader,
} from "@rmwc/drawer";
import { List, ListItem } from "@rmwc/list";

interface Props {}

function Dashboard({}: Props): ReactElement {
  return (
    <div>
       <Router>
      <TopAppBar>
        <TopAppBarRow>
          <TopAppBarSection>
            <TopAppBarTitle>Optimg</TopAppBarTitle>
          </TopAppBarSection>
        </TopAppBarRow>
      </TopAppBar>
      <TopAppBarFixedAdjust />
      <div className="flex">
        <Drawer>
          <DrawerHeader>
            <DrawerTitle>Optimg</DrawerTitle>
          </DrawerHeader>
          <DrawerContent>
            <List>
              <Link to="/optimize"><ListItem>Optimize</ListItem></Link>
              <ListItem>Gallery</ListItem>
              <ListItem>Settings</ListItem>
              <ListItem>Profile</ListItem>
            </List>
          </DrawerContent>
        </Drawer>
        <DrawerAppContent style={{ minHeight: "15rem", padding: "1rem" }}>
         
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="about" element={<About />} />
              <Route path="optimize" element={<Optimize />} />
            </Routes>
         
        </DrawerAppContent>
      </div>
      </Router>
    </div>
  );
}

export default Dashboard;
