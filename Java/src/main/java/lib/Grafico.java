package lib;
import org.jfree.chart.ChartFactory;
import org.jfree.chart.ChartPanel; 
import org.jfree.chart.JFreeChart; 
import org.jfree.chart.plot.PlotOrientation;
import org.jfree.data.category.CategoryDataset; 
import org.jfree.data.category.DefaultCategoryDataset; 
import org.jfree.ui.ApplicationFrame;
import org.jfree.ui.RefineryUtilities; 

public class Grafico extends ApplicationFrame{
        
        private String[] info;
        private String[][] datos;
        
	public Grafico(String[] info, String[][] datos)
	{
            super(info[0]);
            this.info = info;
            this.datos = datos;
		
            JFreeChart graficoBar = ChartFactory.createBarChart(
			info[1], 
			info[2], 
			info[3], 
			crearDataset(datos), 
			PlotOrientation.VERTICAL, 
			true, true, false);
		
            ChartPanel chartPanel = new ChartPanel( graficoBar );        
	    chartPanel.setPreferredSize(new java.awt.Dimension( 560 , 367 ) );        
	    setContentPane( chartPanel );
	}

	private CategoryDataset crearDataset(String[][] datos) 
	{

		final DefaultCategoryDataset dataset = 
			      new DefaultCategoryDataset( );  
		
                for(int i = 0; i < datos.length; i++) {
                    dataset.addValue(Integer.parseInt(datos[i][0]), datos[i][1], datos[i][2]);            
                }
		
		return dataset;
	}
	
        /***
         * Genera el grafico, recibiendo info y datos
         */
	public void Run()
	{
            
            Grafico graf = new Grafico(info, datos);
            graf.pack();
            RefineryUtilities.centerFrameOnScreen( graf );        
	    graf.setVisible( true ); 
	}
	
}
