import { Space, Table, Tag } from 'antd';
import Title from 'antd/lib/typography/Title';
import Column from 'antd/lib/table/Column';
import React, { ReactElement } from 'react';
import Stat from './Stat';
const { ColumnGroup } = Table;

interface Props {}

const data = [
  {
    key: '1',
    rundate: new Date().toLocaleDateString(),
    source: '/mnt',
    destination: '/images',
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '2',
    rundate: new Date().toLocaleDateString(),
    source: '/mnt',
    destination: '/images',
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '3',
    rundate: new Date().toLocaleDateString(),
    source: '/mnt',
    destination: '/images',
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '4',
    rundate: new Date().toLocaleDateString(),
    source: '/mnt',
    destination: '/images',
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
];

export default function Home({}: Props): ReactElement {
    const handleAction = (evt:any) => {
        console.log(evt)
    }
  return (
    <div>
      <div className="flex justify-between">
        {Array.from({ length: 4 }).map(item => {
          return (
            <Stat
              title="Total Images"
              value="1000"
              actions={['JPEG', 'AVIF']}
              onAction={handleAction}
            />
          );
        })}
      </div>
      <section>
        <Title level={3}>All Runs</Title>
        <hr />
        <Table dataSource={data}>
          <Column title="#" dataIndex="key" key="key" />

          <Column title="Date" dataIndex="rundate" key="rundate" />
          <Column title="Source" dataIndex="source" key="source" />
          <Column
            title="Destination"
            dataIndex="destination"
            key="destination"
          />
          <Column
            title="Action"
            key="action"
            render={(text, record: any) => (
              <Space size="middle">
                <a>Invite {record.lastName}</a>
                <a>Delete</a>
              </Space>
            )}
          />
        </Table>
      </section>
    </div>
  );
}
