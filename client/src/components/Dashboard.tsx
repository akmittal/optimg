import React, { ReactElement } from "react";
import Home from "./Home";
import About from "./About";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Optimize from "./Optimize";
import { Typography } from "@rmwc/typography";

import {
  Drawer,
  DrawerTitle,
  DrawerContent,
  DrawerAppContent,
  DrawerHeader,
} from "@rmwc/drawer";
import { List, ListItem } from "@rmwc/list";
import { Server } from "./Server";
import Gallery from "./Gallery";
import GalleryImage from "./GalleryImage";

interface Props {}

function Dashboard({}: Props): ReactElement {
  return (
    <div>
      <Router>
        <div className="flex h-screen">
          <Drawer className="fixed" dismissible open={true}>
            <DrawerHeader>
              <DrawerTitle>
                <img src="/logo.png"  />
              </DrawerTitle>
            </DrawerHeader>
            <DrawerContent>
              <List>
                <Link to="/optimize">
                  <ListItem>Optimize</ListItem>
                </Link>
                <Link to="/gallery">
                  <ListItem>Gallery</ListItem>
                </Link>
                <Link to="/server">
                  <ListItem>Server</ListItem>
                </Link>
              </List>
            </DrawerContent>
          </Drawer>
          <DrawerAppContent className="w-full">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="server" element={<Server />} />
              <Route path="gallery" element={<Gallery />} />
              <Route path="optimize" element={<Optimize />} />
              <Route path="gallery/image" element={<GalleryImage />} />
            </Routes>
          </DrawerAppContent>
        </div>
      </Router>
    </div>
  );
}

export default Dashboard;
