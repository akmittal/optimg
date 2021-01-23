import { Card, Col, Row } from 'antd';
import Meta from 'antd/lib/card/Meta';
import Title from 'antd/lib/typography/Title';
import React, { ReactElement } from 'react';
import { useQuery } from 'react-query';
import { useParams } from 'react-router-dom';
import SizeCompare from '../components/SizeCompare';

interface Props {}

interface ParamTypes {
  path: string;
  name: string;
}

function Image({}: Props): ReactElement {
  const { path = '/', name } = useParams<ParamTypes>();
  const { data, isError, isLoading } = useQuery<any>(
    ['getImagedata', path, name],
    () => {
      const params = new URLSearchParams();
      params.set('path', path);
      params.set('name', name);
      return fetch(`/api/image?${params.toString()}`).then((res: Response) =>
        res.json()
      );
    }
  );
  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error...</div>;
  return (
    <div>
      <div>
        <img
          src={`/api/static/source/${encodeURIComponent(
            path
          )}/${encodeURIComponent(name)}`}
          className="w-6/12 m-auto"
        />
      </div>
      <Title level={3} className="my-4">
        Varients
      </Title>

      <Row>
        {data.varients.map((varient: any) => (
          <Col span={8}>
            <Card
              hoverable
              style={{ width: 240 }}
              cover={
                <img
                  alt="example"
                  src={`/api/static/dest/${encodeURIComponent(varient.path)}/${
                    varient.name
                  }`}
                />
              }
            >
              <Meta
                title={varient.name}
                description={
                 <SizeCompare percentage={getSizePercentage(data.image.size, varient.size)} />
                }
              ></Meta>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}

const getSizePercentage = (baseSize: number, targetSize: number) => {
  return Math.round(((targetSize - baseSize) * 10000) / targetSize) / 100;
};

export default Image;
