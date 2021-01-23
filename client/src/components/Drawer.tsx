import React, { ReactElement } from 'react';

import { Layout, Menu } from 'antd';
import {
  DashboardTwoTone,
  PictureTwoTone,
  SettingTwoTone,
  FundTwoTone,
} from '@ant-design/icons';
import { Link } from 'react-router-dom';

const { SubMenu } = Menu;
const { Sider } = Layout;

interface Props {}

function Drawer({}: Props): ReactElement {
  return (
    <Sider  className="site-layout-background min-h-screen overflow-auto fixed left-0 bottom-0" style={{position:"fixed"}} >
      <Menu 
        mode="inline"
        defaultSelectedKeys={['1']}
        defaultOpenKeys={['sub1']}
        style={{ minHeight: 'calc(100vh)', borderRight: 0 }}
      >
        <img src="/logo.png" />
        <Menu.Item key="1" icon={<DashboardTwoTone />}>
          <Link to="/">Dashboard</Link>
        </Menu.Item>
        <Menu.Item key="2" icon={<PictureTwoTone />}>
          <Link to="/gallery">Gallery</Link>
        </Menu.Item>
        <Menu.Item key="3" icon={<FundTwoTone />}>
          <Link to="/optimize">Optimize</Link>
        </Menu.Item>
        <Menu.Item key="4" icon={<SettingTwoTone />}>
          <Link to="/settings">Settings</Link>
        </Menu.Item>

      </Menu>
    </Sider>
  );
}

export default Drawer;
