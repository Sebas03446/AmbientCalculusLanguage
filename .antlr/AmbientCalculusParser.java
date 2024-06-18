// Generated from /home/jeansebastiansoracipa/GitHub/UNAL/AmbientCalculusLanguage/AmbientCalculus.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class AmbientCalculusParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, AMBIENTID=36, ID=37, INT=38, STRING=39, 
		WS=40, COMMENT=41;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_processDeclaration = 2, RULE_processStatement = 3, 
		RULE_moveStatement = 4, RULE_ambientDeclaration = 5, RULE_conditions = 6, 
		RULE_movementStatement = 7, RULE_sendStatement = 8, RULE_receiveStatement = 9, 
		RULE_inStatement = 10, RULE_outStatement = 11, RULE_openStatement = 12, 
		RULE_callStatement = 13, RULE_printStatement = 14, RULE_printConditionStatement = 15, 
		RULE_modifyConditionStatement = 16, RULE_conditionsBlock = 17, RULE_variableDeclaration = 18, 
		RULE_conditionsVariableDeclaration = 19, RULE_restriction = 20, RULE_assignmentStatement = 21, 
		RULE_expression = 22, RULE_operator = 23, RULE_comparator = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "processDeclaration", "processStatement", "moveStatement", 
			"ambientDeclaration", "conditions", "movementStatement", "sendStatement", 
			"receiveStatement", "inStatement", "outStatement", "openStatement", "callStatement", 
			"printStatement", "printConditionStatement", "modifyConditionStatement", 
			"conditionsBlock", "variableDeclaration", "conditionsVariableDeclaration", 
			"restriction", "assignmentStatement", "expression", "operator", "comparator"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'process'", "'{'", "'}'", "'move'", "';'", "'ambient'", "'to'", 
			"'send'", "'receive'", "'in'", "'out'", "'open'", "'call'", "'print'", 
			"'printc'", "'modifyc'", "'+='", "'-='", "'conditions'", "'let'", "'='", 
			"'letc'", "'restriction'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", 
			"'>'", "'<'", "'>='", "'<='", "'=='", "'!='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "AmbientCalculus.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public AmbientCalculusParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(AmbientCalculusParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 137440100178L) != 0)) {
				{
				{
				setState(50);
				statement();
				}
				}
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(56);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public ProcessDeclarationContext processDeclaration() {
			return getRuleContext(ProcessDeclarationContext.class,0);
		}
		public AmbientDeclarationContext ambientDeclaration() {
			return getRuleContext(AmbientDeclarationContext.class,0);
		}
		public MovementStatementContext movementStatement() {
			return getRuleContext(MovementStatementContext.class,0);
		}
		public SendStatementContext sendStatement() {
			return getRuleContext(SendStatementContext.class,0);
		}
		public ReceiveStatementContext receiveStatement() {
			return getRuleContext(ReceiveStatementContext.class,0);
		}
		public InStatementContext inStatement() {
			return getRuleContext(InStatementContext.class,0);
		}
		public OutStatementContext outStatement() {
			return getRuleContext(OutStatementContext.class,0);
		}
		public OpenStatementContext openStatement() {
			return getRuleContext(OpenStatementContext.class,0);
		}
		public CallStatementContext callStatement() {
			return getRuleContext(CallStatementContext.class,0);
		}
		public PrintStatementContext printStatement() {
			return getRuleContext(PrintStatementContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public ModifyConditionStatementContext modifyConditionStatement() {
			return getRuleContext(ModifyConditionStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(58);
				processDeclaration();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(59);
				ambientDeclaration();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(60);
				movementStatement();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 4);
				{
				setState(61);
				sendStatement();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 5);
				{
				setState(62);
				receiveStatement();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 6);
				{
				setState(63);
				inStatement();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 7);
				{
				setState(64);
				outStatement();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 8);
				{
				setState(65);
				openStatement();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 9);
				{
				setState(66);
				callStatement();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 10);
				{
				setState(67);
				printStatement();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 11);
				{
				setState(68);
				variableDeclaration();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 12);
				{
				setState(69);
				assignmentStatement();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 13);
				{
				setState(70);
				modifyConditionStatement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProcessDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public List<ProcessStatementContext> processStatement() {
			return getRuleContexts(ProcessStatementContext.class);
		}
		public ProcessStatementContext processStatement(int i) {
			return getRuleContext(ProcessStatementContext.class,i);
		}
		public ProcessDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_processDeclaration; }
	}

	public final ProcessDeclarationContext processDeclaration() throws RecognitionException {
		ProcessDeclarationContext _localctx = new ProcessDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_processDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(T__0);
			setState(74);
			match(ID);
			setState(75);
			match(T__1);
			setState(79);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 137440132946L) != 0)) {
				{
				{
				setState(76);
				processStatement();
				}
				}
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(82);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProcessStatementContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public MoveStatementContext moveStatement() {
			return getRuleContext(MoveStatementContext.class,0);
		}
		public PrintConditionStatementContext printConditionStatement() {
			return getRuleContext(PrintConditionStatementContext.class,0);
		}
		public ProcessStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_processStatement; }
	}

	public final ProcessStatementContext processStatement() throws RecognitionException {
		ProcessStatementContext _localctx = new ProcessStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_processStatement);
		try {
			setState(87);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(84);
				statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(85);
				moveStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(86);
				printConditionStatement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MoveStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public MoveStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moveStatement; }
	}

	public final MoveStatementContext moveStatement() throws RecognitionException {
		MoveStatementContext _localctx = new MoveStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_moveStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(T__3);
			setState(90);
			match(AMBIENTID);
			setState(91);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AmbientDeclarationContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public AmbientDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ambientDeclaration; }
	}

	public final AmbientDeclarationContext ambientDeclaration() throws RecognitionException {
		AmbientDeclarationContext _localctx = new AmbientDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ambientDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			match(T__5);
			setState(94);
			match(AMBIENTID);
			setState(95);
			match(T__1);
			setState(97);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__18) {
				{
				setState(96);
				conditions();
				}
			}

			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 137440100178L) != 0)) {
				{
				{
				setState(99);
				statement();
				}
				}
				setState(104);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(105);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsContext extends ParserRuleContext {
		public ConditionsBlockContext conditionsBlock() {
			return getRuleContext(ConditionsBlockContext.class,0);
		}
		public List<ConditionsVariableDeclarationContext> conditionsVariableDeclaration() {
			return getRuleContexts(ConditionsVariableDeclarationContext.class);
		}
		public ConditionsVariableDeclarationContext conditionsVariableDeclaration(int i) {
			return getRuleContext(ConditionsVariableDeclarationContext.class,i);
		}
		public ConditionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditions; }
	}

	public final ConditionsContext conditions() throws RecognitionException {
		ConditionsContext _localctx = new ConditionsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_conditions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			conditionsBlock();
			setState(108);
			match(T__1);
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__21) {
				{
				{
				setState(109);
				conditionsVariableDeclaration();
				}
				}
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(115);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MovementStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public MovementStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_movementStatement; }
	}

	public final MovementStatementContext movementStatement() throws RecognitionException {
		MovementStatementContext _localctx = new MovementStatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_movementStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(T__3);
			setState(118);
			match(ID);
			setState(119);
			match(T__6);
			setState(120);
			match(AMBIENTID);
			setState(121);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SendStatementContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AmbientCalculusParser.STRING, 0); }
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public SendStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sendStatement; }
	}

	public final SendStatementContext sendStatement() throws RecognitionException {
		SendStatementContext _localctx = new SendStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_sendStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(T__7);
			setState(124);
			match(STRING);
			setState(125);
			match(T__6);
			setState(126);
			match(ID);
			setState(127);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ReceiveStatementContext extends ParserRuleContext {
		public ReceiveStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_receiveStatement; }
	}

	public final ReceiveStatementContext receiveStatement() throws RecognitionException {
		ReceiveStatementContext _localctx = new ReceiveStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_receiveStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(T__8);
			setState(130);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public InStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inStatement; }
	}

	public final InStatementContext inStatement() throws RecognitionException {
		InStatementContext _localctx = new InStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_inStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(T__9);
			setState(133);
			match(AMBIENTID);
			setState(134);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OutStatementContext extends ParserRuleContext {
		public OutStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outStatement; }
	}

	public final OutStatementContext outStatement() throws RecognitionException {
		OutStatementContext _localctx = new OutStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_outStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(T__10);
			setState(137);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OpenStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public OpenStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_openStatement; }
	}

	public final OpenStatementContext openStatement() throws RecognitionException {
		OpenStatementContext _localctx = new OpenStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_openStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(T__11);
			setState(140);
			match(AMBIENTID);
			setState(141);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public CallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callStatement; }
	}

	public final CallStatementContext callStatement() throws RecognitionException {
		CallStatementContext _localctx = new CallStatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_callStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(T__12);
			setState(144);
			match(ID);
			setState(145);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printStatement; }
	}

	public final PrintStatementContext printStatement() throws RecognitionException {
		PrintStatementContext _localctx = new PrintStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_printStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(T__13);
			setState(148);
			expression(0);
			setState(149);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintConditionStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public PrintConditionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printConditionStatement; }
	}

	public final PrintConditionStatementContext printConditionStatement() throws RecognitionException {
		PrintConditionStatementContext _localctx = new PrintConditionStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_printConditionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(T__14);
			setState(152);
			match(ID);
			setState(153);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModifyConditionStatementContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ModifyConditionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modifyConditionStatement; }
	}

	public final ModifyConditionStatementContext modifyConditionStatement() throws RecognitionException {
		ModifyConditionStatementContext _localctx = new ModifyConditionStatementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_modifyConditionStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			match(T__15);
			setState(156);
			match(ID);
			setState(157);
			((ModifyConditionStatementContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==T__16 || _la==T__17) ) {
				((ModifyConditionStatementContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(158);
			expression(0);
			setState(159);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsBlockContext extends ParserRuleContext {
		public ConditionsBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionsBlock; }
	}

	public final ConditionsBlockContext conditionsBlock() throws RecognitionException {
		ConditionsBlockContext _localctx = new ConditionsBlockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_conditionsBlock);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_variableDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			match(T__19);
			setState(164);
			match(ID);
			setState(165);
			match(T__20);
			setState(166);
			expression(0);
			setState(167);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsVariableDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public RestrictionContext restriction() {
			return getRuleContext(RestrictionContext.class,0);
		}
		public ConditionsVariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionsVariableDeclaration; }
	}

	public final ConditionsVariableDeclarationContext conditionsVariableDeclaration() throws RecognitionException {
		ConditionsVariableDeclarationContext _localctx = new ConditionsVariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_conditionsVariableDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(T__21);
			setState(170);
			match(ID);
			setState(171);
			match(T__20);
			setState(172);
			match(INT);
			setState(174);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__22) {
				{
				setState(173);
				restriction();
				}
			}

			setState(176);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RestrictionContext extends ParserRuleContext {
		public ComparatorContext comparator() {
			return getRuleContext(ComparatorContext.class,0);
		}
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public RestrictionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_restriction; }
	}

	public final RestrictionContext restriction() throws RecognitionException {
		RestrictionContext _localctx = new RestrictionContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_restriction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(T__22);
			setState(179);
			comparator();
			setState(180);
			match(INT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_assignmentStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(ID);
			setState(183);
			match(T__20);
			setState(184);
			expression(0);
			setState(185);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public TerminalNode STRING() { return getToken(AmbientCalculusParser.STRING, 0); }
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__23:
				{
				setState(188);
				match(T__23);
				setState(189);
				expression(0);
				setState(190);
				match(T__24);
				}
				break;
			case ID:
				{
				setState(192);
				match(ID);
				}
				break;
			case INT:
				{
				setState(193);
				match(INT);
				}
				break;
			case STRING:
				{
				setState(194);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(203);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(197);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(198);
					operator();
					setState(199);
					expression(6);
					}
					} 
				}
				setState(205);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorContext extends ParserRuleContext {
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1006632960L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ComparatorContext extends ParserRuleContext {
		public ComparatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparator; }
	}

	public final ComparatorContext comparator() throws RecognitionException {
		ComparatorContext _localctx = new ComparatorContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_comparator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 67645734912L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 22:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001)\u00d3\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0001\u0000\u0005\u00004\b\u0000\n\u0000\f\u00007\t\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0003\u0001H\b\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0005\u0002N\b\u0002\n\u0002\f\u0002Q\t\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"X\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005b\b\u0005\u0001\u0005"+
		"\u0005\u0005e\b\u0005\n\u0005\f\u0005h\t\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006o\b\u0006\n\u0006\f\u0006"+
		"r\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r"+
		"\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013"+
		"\u00af\b\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u00c4\b\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0005\u0016\u00ca\b\u0016\n\u0016\f\u0016\u00cd"+
		"\t\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0000"+
		"\u0001,\u0019\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \"$&(*,.0\u0000\u0003\u0001\u0000\u0011\u0012"+
		"\u0001\u0000\u001a\u001d\u0001\u0000\u001e#\u00d1\u00005\u0001\u0000\u0000"+
		"\u0000\u0002G\u0001\u0000\u0000\u0000\u0004I\u0001\u0000\u0000\u0000\u0006"+
		"W\u0001\u0000\u0000\u0000\bY\u0001\u0000\u0000\u0000\n]\u0001\u0000\u0000"+
		"\u0000\fk\u0001\u0000\u0000\u0000\u000eu\u0001\u0000\u0000\u0000\u0010"+
		"{\u0001\u0000\u0000\u0000\u0012\u0081\u0001\u0000\u0000\u0000\u0014\u0084"+
		"\u0001\u0000\u0000\u0000\u0016\u0088\u0001\u0000\u0000\u0000\u0018\u008b"+
		"\u0001\u0000\u0000\u0000\u001a\u008f\u0001\u0000\u0000\u0000\u001c\u0093"+
		"\u0001\u0000\u0000\u0000\u001e\u0097\u0001\u0000\u0000\u0000 \u009b\u0001"+
		"\u0000\u0000\u0000\"\u00a1\u0001\u0000\u0000\u0000$\u00a3\u0001\u0000"+
		"\u0000\u0000&\u00a9\u0001\u0000\u0000\u0000(\u00b2\u0001\u0000\u0000\u0000"+
		"*\u00b6\u0001\u0000\u0000\u0000,\u00c3\u0001\u0000\u0000\u0000.\u00ce"+
		"\u0001\u0000\u0000\u00000\u00d0\u0001\u0000\u0000\u000024\u0003\u0002"+
		"\u0001\u000032\u0001\u0000\u0000\u000047\u0001\u0000\u0000\u000053\u0001"+
		"\u0000\u0000\u000056\u0001\u0000\u0000\u000068\u0001\u0000\u0000\u0000"+
		"75\u0001\u0000\u0000\u000089\u0005\u0000\u0000\u00019\u0001\u0001\u0000"+
		"\u0000\u0000:H\u0003\u0004\u0002\u0000;H\u0003\n\u0005\u0000<H\u0003\u000e"+
		"\u0007\u0000=H\u0003\u0010\b\u0000>H\u0003\u0012\t\u0000?H\u0003\u0014"+
		"\n\u0000@H\u0003\u0016\u000b\u0000AH\u0003\u0018\f\u0000BH\u0003\u001a"+
		"\r\u0000CH\u0003\u001c\u000e\u0000DH\u0003$\u0012\u0000EH\u0003*\u0015"+
		"\u0000FH\u0003 \u0010\u0000G:\u0001\u0000\u0000\u0000G;\u0001\u0000\u0000"+
		"\u0000G<\u0001\u0000\u0000\u0000G=\u0001\u0000\u0000\u0000G>\u0001\u0000"+
		"\u0000\u0000G?\u0001\u0000\u0000\u0000G@\u0001\u0000\u0000\u0000GA\u0001"+
		"\u0000\u0000\u0000GB\u0001\u0000\u0000\u0000GC\u0001\u0000\u0000\u0000"+
		"GD\u0001\u0000\u0000\u0000GE\u0001\u0000\u0000\u0000GF\u0001\u0000\u0000"+
		"\u0000H\u0003\u0001\u0000\u0000\u0000IJ\u0005\u0001\u0000\u0000JK\u0005"+
		"%\u0000\u0000KO\u0005\u0002\u0000\u0000LN\u0003\u0006\u0003\u0000ML\u0001"+
		"\u0000\u0000\u0000NQ\u0001\u0000\u0000\u0000OM\u0001\u0000\u0000\u0000"+
		"OP\u0001\u0000\u0000\u0000PR\u0001\u0000\u0000\u0000QO\u0001\u0000\u0000"+
		"\u0000RS\u0005\u0003\u0000\u0000S\u0005\u0001\u0000\u0000\u0000TX\u0003"+
		"\u0002\u0001\u0000UX\u0003\b\u0004\u0000VX\u0003\u001e\u000f\u0000WT\u0001"+
		"\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000WV\u0001\u0000\u0000\u0000"+
		"X\u0007\u0001\u0000\u0000\u0000YZ\u0005\u0004\u0000\u0000Z[\u0005$\u0000"+
		"\u0000[\\\u0005\u0005\u0000\u0000\\\t\u0001\u0000\u0000\u0000]^\u0005"+
		"\u0006\u0000\u0000^_\u0005$\u0000\u0000_a\u0005\u0002\u0000\u0000`b\u0003"+
		"\f\u0006\u0000a`\u0001\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000bf\u0001"+
		"\u0000\u0000\u0000ce\u0003\u0002\u0001\u0000dc\u0001\u0000\u0000\u0000"+
		"eh\u0001\u0000\u0000\u0000fd\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000"+
		"\u0000gi\u0001\u0000\u0000\u0000hf\u0001\u0000\u0000\u0000ij\u0005\u0003"+
		"\u0000\u0000j\u000b\u0001\u0000\u0000\u0000kl\u0003\"\u0011\u0000lp\u0005"+
		"\u0002\u0000\u0000mo\u0003&\u0013\u0000nm\u0001\u0000\u0000\u0000or\u0001"+
		"\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000"+
		"qs\u0001\u0000\u0000\u0000rp\u0001\u0000\u0000\u0000st\u0005\u0003\u0000"+
		"\u0000t\r\u0001\u0000\u0000\u0000uv\u0005\u0004\u0000\u0000vw\u0005%\u0000"+
		"\u0000wx\u0005\u0007\u0000\u0000xy\u0005$\u0000\u0000yz\u0005\u0005\u0000"+
		"\u0000z\u000f\u0001\u0000\u0000\u0000{|\u0005\b\u0000\u0000|}\u0005\'"+
		"\u0000\u0000}~\u0005\u0007\u0000\u0000~\u007f\u0005%\u0000\u0000\u007f"+
		"\u0080\u0005\u0005\u0000\u0000\u0080\u0011\u0001\u0000\u0000\u0000\u0081"+
		"\u0082\u0005\t\u0000\u0000\u0082\u0083\u0005\u0005\u0000\u0000\u0083\u0013"+
		"\u0001\u0000\u0000\u0000\u0084\u0085\u0005\n\u0000\u0000\u0085\u0086\u0005"+
		"$\u0000\u0000\u0086\u0087\u0005\u0005\u0000\u0000\u0087\u0015\u0001\u0000"+
		"\u0000\u0000\u0088\u0089\u0005\u000b\u0000\u0000\u0089\u008a\u0005\u0005"+
		"\u0000\u0000\u008a\u0017\u0001\u0000\u0000\u0000\u008b\u008c\u0005\f\u0000"+
		"\u0000\u008c\u008d\u0005$\u0000\u0000\u008d\u008e\u0005\u0005\u0000\u0000"+
		"\u008e\u0019\u0001\u0000\u0000\u0000\u008f\u0090\u0005\r\u0000\u0000\u0090"+
		"\u0091\u0005%\u0000\u0000\u0091\u0092\u0005\u0005\u0000\u0000\u0092\u001b"+
		"\u0001\u0000\u0000\u0000\u0093\u0094\u0005\u000e\u0000\u0000\u0094\u0095"+
		"\u0003,\u0016\u0000\u0095\u0096\u0005\u0005\u0000\u0000\u0096\u001d\u0001"+
		"\u0000\u0000\u0000\u0097\u0098\u0005\u000f\u0000\u0000\u0098\u0099\u0005"+
		"%\u0000\u0000\u0099\u009a\u0005\u0005\u0000\u0000\u009a\u001f\u0001\u0000"+
		"\u0000\u0000\u009b\u009c\u0005\u0010\u0000\u0000\u009c\u009d\u0005%\u0000"+
		"\u0000\u009d\u009e\u0007\u0000\u0000\u0000\u009e\u009f\u0003,\u0016\u0000"+
		"\u009f\u00a0\u0005\u0005\u0000\u0000\u00a0!\u0001\u0000\u0000\u0000\u00a1"+
		"\u00a2\u0005\u0013\u0000\u0000\u00a2#\u0001\u0000\u0000\u0000\u00a3\u00a4"+
		"\u0005\u0014\u0000\u0000\u00a4\u00a5\u0005%\u0000\u0000\u00a5\u00a6\u0005"+
		"\u0015\u0000\u0000\u00a6\u00a7\u0003,\u0016\u0000\u00a7\u00a8\u0005\u0005"+
		"\u0000\u0000\u00a8%\u0001\u0000\u0000\u0000\u00a9\u00aa\u0005\u0016\u0000"+
		"\u0000\u00aa\u00ab\u0005%\u0000\u0000\u00ab\u00ac\u0005\u0015\u0000\u0000"+
		"\u00ac\u00ae\u0005&\u0000\u0000\u00ad\u00af\u0003(\u0014\u0000\u00ae\u00ad"+
		"\u0001\u0000\u0000\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b0"+
		"\u0001\u0000\u0000\u0000\u00b0\u00b1\u0005\u0005\u0000\u0000\u00b1\'\u0001"+
		"\u0000\u0000\u0000\u00b2\u00b3\u0005\u0017\u0000\u0000\u00b3\u00b4\u0003"+
		"0\u0018\u0000\u00b4\u00b5\u0005&\u0000\u0000\u00b5)\u0001\u0000\u0000"+
		"\u0000\u00b6\u00b7\u0005%\u0000\u0000\u00b7\u00b8\u0005\u0015\u0000\u0000"+
		"\u00b8\u00b9\u0003,\u0016\u0000\u00b9\u00ba\u0005\u0005\u0000\u0000\u00ba"+
		"+\u0001\u0000\u0000\u0000\u00bb\u00bc\u0006\u0016\uffff\uffff\u0000\u00bc"+
		"\u00bd\u0005\u0018\u0000\u0000\u00bd\u00be\u0003,\u0016\u0000\u00be\u00bf"+
		"\u0005\u0019\u0000\u0000\u00bf\u00c4\u0001\u0000\u0000\u0000\u00c0\u00c4"+
		"\u0005%\u0000\u0000\u00c1\u00c4\u0005&\u0000\u0000\u00c2\u00c4\u0005\'"+
		"\u0000\u0000\u00c3\u00bb\u0001\u0000\u0000\u0000\u00c3\u00c0\u0001\u0000"+
		"\u0000\u0000\u00c3\u00c1\u0001\u0000\u0000\u0000\u00c3\u00c2\u0001\u0000"+
		"\u0000\u0000\u00c4\u00cb\u0001\u0000\u0000\u0000\u00c5\u00c6\n\u0005\u0000"+
		"\u0000\u00c6\u00c7\u0003.\u0017\u0000\u00c7\u00c8\u0003,\u0016\u0006\u00c8"+
		"\u00ca\u0001\u0000\u0000\u0000\u00c9\u00c5\u0001\u0000\u0000\u0000\u00ca"+
		"\u00cd\u0001\u0000\u0000\u0000\u00cb\u00c9\u0001\u0000\u0000\u0000\u00cb"+
		"\u00cc\u0001\u0000\u0000\u0000\u00cc-\u0001\u0000\u0000\u0000\u00cd\u00cb"+
		"\u0001\u0000\u0000\u0000\u00ce\u00cf\u0007\u0001\u0000\u0000\u00cf/\u0001"+
		"\u0000\u0000\u0000\u00d0\u00d1\u0007\u0002\u0000\u0000\u00d11\u0001\u0000"+
		"\u0000\u0000\n5GOWafp\u00ae\u00c3\u00cb";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}