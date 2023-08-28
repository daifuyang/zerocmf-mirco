import Guide from '@/components/Guide';
import { useModel } from '@umijs/max';
import styles from './index.less';

const HomePage: React.FC = () => {
  const { name } = useModel('global');
  return (
    <>
       <div className={styles.container}>
          <Guide name={name} />
        </div>
    </>
  );
};

export default HomePage;
