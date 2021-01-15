import { Card, Col, Row } from 'antd';
import Meta from 'antd/lib/card/Meta';
import Title from 'antd/lib/typography/Title';
import React, { ReactElement } from 'react';
import { useParams } from 'react-router-dom';

interface Props {}

interface ParamTypes {
  path: string;
  name: string;
}

function Image({}: Props): ReactElement {
  const {path = "/", name} = useParams<ParamTypes>()
  return (
    <div>
      <div>
        <img
          src={`/api/static/source${path}/${name}`}
          className="w-6/12 m-auto"
        />
      </div>
      <Title level={3} className="my-4">Varients</Title>
     
      <Row >
        {[1, 2, 3].map(_ => (
          <Col span={8}>
            <Card
            
              hoverable
              style={{ width: 240 }}
              cover={
                <img
                  alt="example"
                  src={`/api/static/source`}
                />
              }
            >
              <Meta
                title="Some Name"
                description={<div>helllllo</div>}
              
              ></Meta>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}

export default Image;
