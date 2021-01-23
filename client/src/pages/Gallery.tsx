import { Card, Col, Pagination, Row } from 'antd';
import Meta from 'antd/lib/card/Meta';
import React, { ReactElement } from 'react';
import { useQuery } from 'react-query';
import { Spin } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';
import { Link, useHistory, useParams } from 'react-router-dom';
import { LinkOutlined } from '@ant-design/icons';

interface ParamTypes {
  pageNo: string;
  path: string;
}

interface Props {}
const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin />;

const PageSize = 15;

function Gallery({}: Props): ReactElement {
  const history = useHistory();
  const { pageNo = '1', path = '/' } = useParams<ParamTypes>();
  const { data, isError, isLoading } = useQuery(
    ['galleryData', pageNo, path],
    () => {
      const params = new URLSearchParams();
      params.set('page', pageNo);
      params.set('path', path);
      return fetch(`/api/gallery?${params.toString()}`).then(res => res.json());
    }
  );
  return (
    <>
      {isLoading && <Spin indicator={antIcon} />}
      {!isLoading && !isError && (
        <div>
          <Row gutter={[16, 16]}>
            {data.images && data.images.map((image: any) => (
              <Col span={8}>
                <Card
                  hoverable
                  style={{ width: 240 }}
                  actions={[
                    <Link
                      to={`/image/${encodeURIComponent(
                        image.image.path
                      )}/${encodeURIComponent(image.image.name)}`}
                    >
                      <LinkOutlined key="open" />
                    </Link>,
                  ]}
                  cover={
                    <img
                      alt="example"
                      src={`/api/static/source/${image.image.path}/${image.image.name}`}
                    />
                  }
                >
                  <Meta
                    title={image.image.name}
                    description={
                      <section>
                        <div>
                          Size: {Math.floor(image.image.size / 1024)} KB
                        </div>
                        <div>Varients: {image.varients.length}</div>
                      </section>
                    }
                  />
                </Card>
              </Col>
            ))}
          </Row>

          <div className="flex justify-end my-4">
            <Pagination
              defaultCurrent={data.currentPage}
              total={data.totalPages * PageSize}
              defaultPageSize={PageSize}
              showSizeChanger={false}
              onChange={(page: number) => {
                history.push(`/gallery/${encodeURIComponent(path)}/${page}`);
              }}
            />
          </div>
        </div>
      )}
    </>
  );
}

export default Gallery;
