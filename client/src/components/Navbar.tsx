import React, { ReactElement } from 'react';
import { Breadcrumb, Layout, Menu, PageHeader } from 'antd';

const { Header } = Layout;
interface Route {
  path: string;
  breadcrumbName: string;
}

interface Props {
  title: string;
  routes: Route[];
}

function Navbar({ title, routes }: Props): ReactElement {
  return (
    <PageHeader
      className="site-page-header"
      title={title}
      subTitle="|"
      footer={
        <Breadcrumb>
          { routes && routes.map((route: Route) => (
            <Breadcrumb.Item>{route.breadcrumbName}</Breadcrumb.Item>
          ))}
        </Breadcrumb>
      }
    />
  );
}

export default Navbar;
