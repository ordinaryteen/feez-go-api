import React, { useEffect, useState } from 'react';
import { Box, SimpleGrid, Container, Text, Spinner, Center } from '@chakra-ui/react';
import ProductCard from '../components/ProductCard'; 

function HomePage() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/products`);
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error("Gagal fetch produk:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchProducts();
  }, []);

  return (
    <Container maxW="1200px" py={8}>
      <Text fontSize="2xl" fontWeight="bold" color="white" mb={6}>
        Latest Drops
      </Text>

      {loading ? (
        <Center h="50vh">
          <Spinner size="xl" color="brand.500" />
        </Center>
      ) : (
        <SimpleGrid columns={[1, 2, 3, 4]} spacing={6}>
          {products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </SimpleGrid>
      )}
    </Container>
  );
}

export default HomePage;