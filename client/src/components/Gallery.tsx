import React, { ReactElement, useEffect, useState } from "react";
import { useQuery } from "react-query";

import Header from "./Header";
import ImageCard from "./ImageCard";
import ReactPaginate from "react-paginate";

interface Props {}

export default function Gallery(props: Props): ReactElement {
  const [currentPage, setCurrentPage] = useState(0);
  const { isLoading, error, data, refetch } = useQuery("repoData", () =>
    fetch(`/gallery?page=${currentPage}`).then((res) => res.json())
  );
  useEffect(() => {
    refetch();
  }, [currentPage]);

  if (isLoading) return <div>Loading...</div>;

  if (error) return <div>An error has occurred: ' + error.message</div>;

  return (
    <div>
      <Header text="Gallery" />
      {/* {JSON.stringify(data)} */}
      <section className="flex flex-row flex-wrap justify-around">
        {data.Data.map((item: any) => (
          <ImageCard
            name={item.Main.Path}
            imageURL={item.Main.Path}
            height={item.Main.Height}
            width={item.Main.Width}
            size={item.Main.Size}
            modified={item.Main.Modified}
          />
        ))}
      </section>
      <ReactPaginate
        pageCount={data.TotalPages}
        pageRangeDisplayed={4}
        marginPagesDisplayed={2}
        containerClassName={"flex justify-end m-4"}
        previousLabel={"‹"}
        nextLabel={"›"}
        pageClassName="p-2 "
        previousClassName="p-2 text-xl"
        nextClassName="p-2 text-xl"
        activeClassName={"bg-gray-200"}
        onPageChange={(item) => {
          setCurrentPage(item.selected);
        }}
      />
    </div>
  );
}
